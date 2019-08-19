package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ning-hu/mh-line-bot/linebot"
)

var bot *linebot.Client

LINE_ID := map[string]string{
	"Ning": "U275f7f23c237a5589177d1d32830389",
}

func main() {
	secret := os.Getenv("LINE_SECRET")
	accessToken := os.Getenv("LINE_ACCESS_TOKEN")

	var err error
	bot, err = linebot.New(secret, accessToken)
	if err != nil {
		log.Fatal("Error creating a new http handler")
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Println("Invalid Signature")
				w.WriteHeader(400)
			} else {
				log.Println("Server Error")
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			log.Println(event.Source.UserID)
			if event.Type == linebot.EventTypeMessage {
				fmt.Printf("%+v\n", event)
				fmt.Printf("%+v\n", event.Source)
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if strings.Contains(message.Text, "@Ning") && strings.Contains(message.Text, "?") {
						sendMessage(event.ReplyToken, `¯\_(ツ)_/¯`)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func sendMessage(token, message string) {
	if _, err := bot.ReplyMessage(token, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
	}
}
