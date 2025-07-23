package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const tenMb int64 = 10 * 1024 * 1024

var storeCmd = &cobra.Command{
	Use:   "store [file]",
	Short: "Store an encrypted backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		inputFilePath := fmt.Sprintf("src/%s", file)
		destinationPath := fmt.Sprintf("src/%s.bak", file)
		fileSize := getFileSize(inputFilePath)

		if fileSize < tenMb {
			finput, err := os.ReadFile(inputFilePath)
			if err != nil {
				defineError(err, "Error reading file")
			}
			if err := os.WriteFile(destinationPath, finput, 0644); err != nil {
				defineError(err, "Error writing file")
			}

		} else {
			finput, err := os.Open(inputFilePath)

			if err != nil {
				defineError(err, "Unable to open input file")
			}

			defer func() {
				if err := finput.Close(); err != nil {
					defineError(err, "Unable to close input file")
				}
			}()
			fout, err := os.Create(destinationPath)

			if err != nil {
				defineError(err, "Unable to open output file")
			}

			defer func() {
				if err := fout.Close(); err != nil {
					defineError(err, "Unable to close output file")
				}
			}()

			if _, err := io.Copy(fout, finput); err != nil {
				defineError(err, "Error copying file")
			}
		}

		return nil
	},
}
