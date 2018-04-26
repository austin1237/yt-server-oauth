package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
)

type YtAuth struct {
	ClientID     string    `json:"clientID"`
	ClientSecret string    `json:"clientSecret"`
	AccessToken  string    `json:"accessToken"`
	TokenType    string    `json:"tokenType"`
	RefreshToken string    `json:"refreshToken"`
	Expiry       time.Time `json:"expiry"`
}

var (
	//ENV VARIABLES
	youtubeClientID string
	youtubeSecret   string

	scopes []string
	//You can find a full list of the scope at https://developers.google.com/identity/protocols/googlescopes#youtubev3
)

func init() {
	youtubeClientID = os.Getenv("YOUTUBE_CLIENT_ID")
	if youtubeClientID == "" {
		fmt.Println("YOUTUBE_CLIENT_ID ENV var was not set.")
		os.Exit(1)
	}

	youtubeSecret = os.Getenv("YOUTUBE_SECRET")
	if youtubeSecret == "" {
		fmt.Println("YOUTUBE_SECRET ENV var was not set.")
		os.Exit(1)
	}

	//You can find a full list of the avaible scopes for youtube here
	//https://developers.google.com/identity/protocols/googlescopes#youtubev3

	scopes = []string{"https://www.googleapis.com/auth/youtube.upload"}
}

func logAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting due to the above error")
		os.Exit(1)
	}
}

func main() {
	config := generateConfig(scopes)
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	tok, err := getTokenFromPrompt(config, authURL)
	if err != nil {
		logAndExit(err)
	}
	encoded, err := generateBase64String(config, tok)
	if err != nil {
		logAndExit(err)
	}
	fmt.Println("The following is your base64 youtube creds:")
	fmt.Println(encoded)
}

func generateConfig(scopes []string) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     youtubeClientID,
		ClientSecret: youtubeSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
	return config
}

func generateBase64String(config *oauth2.Config, tok *oauth2.Token) (string, error) {
	ytAuth := YtAuth{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		AccessToken:  tok.AccessToken,
		TokenType:    tok.TokenType,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry,
	}
	jsonBytes, err := json.Marshal(ytAuth)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)
	return encoded, nil
}

// Exchange the authorization code for an access token
func exchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token %v", err)
		return tok, err
	}
	return tok, nil
}

// getTokenFromPrompt uses Config to request a Token and prompts the user
// to enter the token on the command line. It returns the retrieved Token.
func getTokenFromPrompt(config *oauth2.Config, authURL string) (*oauth2.Token, error) {
	var code string
	fmt.Printf("Go to the following link in your browser. After completing "+
		"the authorization flow, enter the authorization code on the command "+
		"line: \n%v\n", authURL)

	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
		return &oauth2.Token{}, err
	}
	return exchangeToken(config, code)
}
