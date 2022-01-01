package main

import (
	"fmt"
	"log"

	"github.com/abhinavxd/twitch-live-chat-downloader"
)

// twitch channel name from URL-  https://www.twitch.tv/aceu
const CHANNEL_NAME = "aceu"

func main() {
	// Open a websocket connection to twitch chat
	// Call this function again to reconnect
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

		// Optionally you can Parse tags which returns a Message struct
		// Or you can use the raw twitch WS tags
		parsedMsg, err := twitch_chat.ParseTags(message)
		
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("%+v\n\n", parsedMsg)
	}
}