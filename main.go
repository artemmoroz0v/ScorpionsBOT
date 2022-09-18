package main

import (
	"fmt"
	"log"
	f "ScorpionsBOT/functions"
)

func main() {
	var bot_token string
	fmt.Print("Please enter telegram-bot-token!\n")
	fmt.Scan(&bot_token)
	bot_API := "https://api.telegram.org/bot"
	bot_URL := bot_API + bot_token
	offset := 0
	for {
		updates, err := f.GetUpdates(bot_URL, offset)
		if err != nil {
			log.Println("An error has been detected: ", err.Error())
		}
		for _, update := range updates {
			offset = update.UpdateID + 1
			err = f.Answer(bot_URL, update)
		}
		fmt.Print(updates, "\n")
	}
}
