package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ning-hu/mh-line-bot/linebot"
	"github.com/ning-hu/mh-line-bot/linebot/httphandler"
)

func main() {
	secret := os.Getenv("LINE_SECRET")
	accessToken := os.Getenv("LINE_ACESS_TOKEN")

	handler, err := httphandler.New(secret, accessToken)
	if err != nil {
		log.Fatal("Error creating a new http handler")
	}

	handlerFunc := func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", events)

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("hello")).Do()
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			}
		}
	}
	handler.HandleEvents(handlerFunc)
}
