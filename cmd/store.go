package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"log"
	"os"
	"path/filepath"
	"syno/misc"
	storageconfig "syno/storageConfig"
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
		file, err := os.Open(file_Path)
		baseFilename := filepath.Base(file_Path)
		if err != nil {
			misc.DefineError(err, "Unable to open file")
			return err
		}
		defer file.Close()

		driveFile, err := server.Files.Create(&drive.File{Name: baseFilename}).Media(file).Do()

		if err != nil {
			misc.DefineError(err, "Unable to upload file to google drive")
			return err
		}
		log.Printf("File successfully uploaded! ID: %s, Name: %s\n", driveFile.Id, driveFile.Name)
		return nil
	},
}
