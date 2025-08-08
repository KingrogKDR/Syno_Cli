package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "syno",
	Short: "A secure CLI sync tool for cloud storage",
	Long:  `syno is a tool that securely syncs files to any kind of storage you like`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(storeCmd)
	rootCmd.AddCommand(retrieveCmd)
}
