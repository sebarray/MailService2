package mail

import (
	"context"
	"fmt"
	"log"
	"mailgo/model"
	"os"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var GmailService *gmail.Service

func OAuthGmailService() {
	wg := &sync.WaitGroup{}
	var token model.Token
	rtoken := os.Getenv("RTOKEN")
	config := oauth2.Config{
		ClientID:     os.Getenv("CID"),
		ClientSecret: os.Getenv("CSECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       []string{},
	}
	wg.Add(1)
	go UpdateTokenAcces(wg, &token)
	wg.Wait()
	tokens := oauth2.Token{
		AccessToken:  token.Access,
		TokenType:    "Bearer",
		RefreshToken: rtoken,
		Expiry:       time.Time{},
	}

	var tokenSource = config.TokenSource(context.Background(), &tokens)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is initialized ")
	}

}
