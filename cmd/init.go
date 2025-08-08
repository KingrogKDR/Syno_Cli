package cmd

import (
	"fmt"
	storageconfig "syno/storageConfig"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and authenticate with Google Drive",
	Long:  `This command will guide you through the process of authenticating with your Google Drive account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing Google Drive authentication...")
		fmt.Println("1. Ensure you have a 'credentials.json' file from the Google Cloud Console.")
		fmt.Println("2. This command will open a browser window for you to authorize access. If authorized already, it will refresh the token.json")

		storageconfig.GetClient(storageconfig.GetConfig())

		fmt.Println("\nAuthentication successful! Your token has been saved.")
	},
}
