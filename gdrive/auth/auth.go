package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alon7/gdrive/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Get uath code from user
func promptUserForAuthCode(config *oauth2.Config) string {
	authUrl := config.AuthCodeURL("state")
	fmt.Println("Go to the following link in your browser:")
	fmt.Printf("%v\n\n", authUrl)
	return util.Prompt("Enter verification code: ")
}
