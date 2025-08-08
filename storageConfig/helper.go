package storageconfig

import (
	"context"
	"net/http"
	"os"
	"syno/misc"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GetClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}
func GetConfig() *oauth2.Config {
	buf, err := os.ReadFile(misc.PathToCredentials)
	if err != nil {
		misc.DefineError(err, "Unable to read credentials file")
		return nil
	}
	config, err := google.ConfigFromJSON(buf, drive.DriveScope)

	if err != nil {
		misc.DefineError(err, "Unable to get driver configuration")
		return nil
	}

	return config
}

func GetDriveService() (*drive.Service, error) {
	client := GetClient(GetConfig())
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		misc.DefineError(err, "Unable to get drive Service")
		return nil, err
	}
	return srv, err
}
