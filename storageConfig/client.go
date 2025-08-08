package storageconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"syno/misc"

	"golang.org/x/oauth2"
)

func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authUrl := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authUrl)
	var authCode string
	fmt.Print("Auth Code: ")
	if _, err := fmt.Scan(&authCode); err != nil {
		misc.DefineError(err, "Unable to read authorization code")
	}
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		misc.DefineError(err, "Unable to retrieve token from web")
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		misc.DefineError(err, "Unable to cache oauth token")
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
