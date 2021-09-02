package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Realizar um tweet
	tweet, _, err := client.Statuses.Update("Uma mensagem de teste", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(tweet.Text)

	// Procurar por tweets
	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#reactnative",
		Count: 2,
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, val := range tweets.Statuses {
		log.Print("Nome do usuário: ", val.User.Name)
		log.Print("Tweet: ", val.Text)

		_, _, err := client.Statuses.Retweet(val.ID, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
