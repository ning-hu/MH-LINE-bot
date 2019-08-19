package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ning-hu/mh-line-bot/linebot"
)

var bot *linebot.Client

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
					resMessage := []string{}
					if os.Getenv("USER") == "rick" {
						if strings.Contains(message.Text, "@Rick") {
							// emoji, _ := strconv.ParseInt("\uDBC0\uDC84", 0, 16)
							// emojiString := fmt.Sprintf("I'm busy. Ask Lev \uDBC0\uDC84", string(emoji))
							resMessage = append(resMessage, "I'm busy. Ask Lev \uDBC0\uDC84")
						}
					} else {
						if strings.Contains(message.Text, "@Ning") && strings.Contains(message.Text, "?") {
							resMessage = append(resMessage, `¯\_(ツ)_/¯`)
						}
					}
					sendMessage(event.ReplyToken, strings.Join(resMessage, " "))
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
