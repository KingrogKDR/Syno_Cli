package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syno/misc"
	storageconfig "syno/storageConfig"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

var storeCmd = &cobra.Command{
	Use:   "push [file]",
	Short: "Uploads a file to Google Drive",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Uploading ....")
		server, err := storageconfig.GetDriveService()
		if err != nil {
			misc.DefineError(err, "Unable to retrieve Drive client")
			return err
		}
		file_Path := args[0]
		baseFilename := filepath.Base(file_Path)

		query := fmt.Sprintf("name = '%s' and trashed = false", baseFilename)
		r, err := server.Files.List().Q(query).Fields("files(id, name)").Do()
		if err != nil {
			misc.DefineError(err, "Unable to search for existing files")
			return err
		}

		if len(r.Files) > 0 {
			existingFile := r.Files[0]
			fmt.Printf("A file named '%s' already exists (ID: %s).\n", existingFile.Name, existingFile.Id)
			fmt.Print("Do you want to [r]eplace it or [c]reate a new file? (r/c): ")

			var choice string
			_, err := fmt.Scanln(&choice)
			if err != nil || (choice != "r" && choice != "R" && choice != "c" && choice != "C") {
				fmt.Println("Invalid choice. Aborting upload.")
				return nil
			}

			if choice == "r" || choice == "R" {
				return replaceFile(file_Path, server, baseFilename, existingFile.Id)
			} else {
				return createUniqueFile(server, baseFilename, file_Path)
			}
		}
		return createNewFile(file_Path, server, baseFilename)
	},
}

func createNewFile(file_Path string, server *drive.Service, baseFilename string) error {
	file, err := os.Open(file_Path)
	if err != nil {
		misc.DefineError(err, "Unable to open file in createNewFile")
		return err
	}
	defer file.Close()

	driveFile, err := server.Files.Create(&drive.File{Name: baseFilename}).Media(file).Do()

	if err != nil {
		misc.DefineError(err, "Unable to upload file to Google Drive")
		return err
	}

	log.Printf("New file successfully uploaded! ID: %s, Name: %s\n", driveFile.Id, driveFile.Name)

	return nil
}

func replaceFile(file_Path string, server *drive.Service, baseFilename string, fileID string) error {
	file, err := os.Open(file_Path)
	if err != nil {
		misc.DefineError(err, "Unable to open file in createNewFile")
		return err
	}
	defer file.Close()

	driveFile, err := server.Files.Update(fileID, &drive.File{Name: baseFilename}).Media(file).Do()
	if err != nil {
		misc.DefineError(err, "Unable to replace file in Google Drive")
		return err
	}

	log.Printf("File '%s' successfully replaced! ID: %s\n", driveFile.Name, driveFile.Id)
	return nil
}

func createUniqueFile(server *drive.Service, baseFilename string, file_Path string) error {
	ext := filepath.Ext(baseFilename)
	nameWithoutExt := baseFilename[:len(baseFilename)-len(ext)]
	i := 1
	var uniqueFilename string

	for {
		uniqueFilename = fmt.Sprintf("%s (%d)%s", nameWithoutExt, i, ext)
		query := fmt.Sprintf("name = '%s' and trashed = false", uniqueFilename)
		r, err := server.Files.List().Q(query).Fields("files(id)").Do()
		if err != nil {
			misc.DefineError(err, "Unable to check for unique filename")
			return err
		}

		if len(r.Files) == 0 {
			// Found a unique name, break the loop
			break
		}

		i++
	}
	file, err := os.Open(file_Path)
	if err != nil {
		misc.DefineError(err, "Unable to open file in createUniqueFile")
		return err
	}
	defer file.Close()

	driveFile, err := server.Files.Create(&drive.File{Name: uniqueFilename}).Media(file).Do()
	if err != nil {
		misc.DefineError(err, "Unable to upload file to Google Drive with a unique name")
		return err
	}

	log.Printf("New file successfully uploaded! ID: %s, Name: %s\n", driveFile.Id, driveFile.Name)
	return nil
}
