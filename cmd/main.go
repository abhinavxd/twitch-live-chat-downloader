package main

import (
	"fmt"
	"log"
	"os"

	twitch_chat "github.com/abhinavxd/twitch-live-chat-downloader"

	flag "github.com/spf13/pflag"
)

func main() {
	// Parse command line flags
	flagSet := flag.NewFlagSet("twitch-chat", flag.ContinueOnError)

	// Parse commandline flags
	flagSet.String("channel", "aceu", "Name of the channel to fetch chat from")
	flagSet.Bool("verbose", false, "Verbose mode prints meta information")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	channel, err := flagSet.GetString("channel")
	if err != nil {
		log.Fatalf("error parsing channel flag: %v", err)
	}

	vb, err := flagSet.GetBool("verbose")
	if err != nil {
		log.Fatalf("error parsing verbose flag: %v", err)
	}

	if err = twitch_chat.InitializeConnection(channel); err != nil {
		log.Fatal(err)
	}

	for {
		message, err := twitch_chat.FetchMessages()
		if err != nil {
			log.Println(err)
			// Reconnect websocket connection
			if err := twitch_chat.InitializeConnection(channel); err != nil {
				log.Println(err)
			}
			continue
		}

		// Optionally you can Parse IRC tags which returns a Message struct
		// Or you can use the IRC tags directly
		parsedMsg, err := twitch_chat.ParseTags(message)
		if err != nil {
			log.Println(err)
			continue
		}

		if !vb {
			fmt.Printf("%s : %s\n", parsedMsg.Username, parsedMsg.Message)
		} else {
			fmt.Printf("%+v\n", parsedMsg)
		}
	}
}

