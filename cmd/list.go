package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"google.golang.org/api/googleapi"
	"os"
	"strconv"
	"syno/misc"
	storageconfig "syno/storageConfig"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists files in your Google Drive",
	Long:  `This command lists all the files in your Google Drive, including their ID, name, and size.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := storageconfig.GetDriveService()
		if err != nil {
			misc.DefineError(err, "Unable to retrieve Drive client")
			return err
		}
		query := "trashed=false"
		fields := "nextPageToken, files(id, name, size, mimeType)"

		res, err := server.Files.List().Q(query).Fields(googleapi.Field(fields)).Do()

		if err != nil {
			misc.DefineError(err, "Unable to list files")
			return err
		}
		if len(res.Files) == 0 {
			fmt.Println("No files found.")
			return nil
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"ID", "NAME", "SIZE", "TYPE"})

		termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Could not determine terminal width:", err)
			termWidth = 100
		}
		maxNameWidth := max(10, termWidth-(10+10+20+16+4*3))
		var data [][]string
		for _, i := range res.Files {
			name := truncateString(i.Name, maxNameWidth)
			size := formatSize(i.Size)

			row := []string{i.Id, name, size, i.MimeType}
			data = append(data, row)
		}

		table.Bulk(data)

		table.Render()
		return nil
	},
}

func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

func formatSize(s int64) string {
	if s < 1024 {
		return strconv.FormatInt(s, 10) + " B"
	}
	f := float64(s)
	for _, unit := range []string{"KB", "MB", "GB", "TB"} {
		f /= 1024
		if f < 1024 {
			return fmt.Sprintf("%.2f %s", f, unit)
		}
	}
	return fmt.Sprintf("%.2f %s", f, "PB")
}
