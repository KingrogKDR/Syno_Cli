package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "pull [file]",
	Short: "Retrieve the backup and decrypts it, storing it at the preferred location",
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		fmt.Printf("%s\n", file)

		return nil
	},
}
