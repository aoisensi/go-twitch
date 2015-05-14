package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aoisensi/go-twitch/v3"
	"github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

var scopes = []string{"user_read",
	"user_blocks_edit",
	"user_blocks_read",
	"user_follows_edit",
	"channel_read",
	"channel_editor",
	"channel_commercial",
	"channel_stream",
	"channel_subscriptions",
	"user_subscriptions",
	"channel_check_subscription",
	"chat_login",
}

func main() {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	clientSecret := os.Getenv("TWITCH_CLIENT_SECRET")
	redirectUrl := os.Getenv("TWITCH_REDIRECT_URL")
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.twitch.tv/kraken/oauth2/authorize",
			TokenURL: "https://api.twitch.tv/kraken/oauth2/token",
		},
		RedirectURL: redirectUrl,
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println("Visit the URL for the auth dialog:")
	fmt.Println(url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}
	token.TokenType = "OAuth"

	client := conf.Client(oauth2.NoContext, token)
	twi, _ := twitch.NewClient(client)
	out, err := twi.Channels.GetSelf()
	pp.Println(out)
	pp.Println(err)
}
