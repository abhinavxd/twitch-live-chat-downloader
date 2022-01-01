package main

import (
	"fmt"
	"log"

	"github.com/abhinavxd/twitch-live-chat-downloader"
)

// twitch channel name from URL-  https://www.twitch.tv/aceu
const CHANNEL_NAME = "aceu"

func main() {
	err := twitch_chat.InitializeConnection(CHANNEL_NAME)
	if err != nil {
		log.Fatal(err)
	}
	for {
		message, err := twitch_chat.FetchMessages()
		
		if err != nil {
			log.Println(err)
			continue
		}

		// Optionally you can Parse messages which returns a Message struct
		parsedMsg, err := twitch_chat.ParseTags(message)
		
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("%+v\n", parsedMsg)
	}
}