package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store [file]",
	Short: "Store an encrypted backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		filePath := "src/" + file

		f, err := os.Open(filePath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open file: %v", err)
		}

		// open and read the file
		// copy the contents of the file and store it in a destination at syno in a folder called backup
		// close thefiles

		return nil
	},
}
