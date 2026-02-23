package misc

import (
	"os"
	"path/filepath"
)

var PathToCredentials string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		DefineError(err, "Unable to get user home directory")
	}
	PathToCredentials = filepath.Join(homeDir, "Documents", "imp", "credentials.json")
}
