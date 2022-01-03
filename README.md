# twitch-live-chat-downloader

![pepe](https://user-images.githubusercontent.com/48166553/147852858-7ecc4ece-ebcb-4b75-a6a2-752bbeb896fa.gif) Twitch chat feed without any authentication


## How does it work?
* Simply mimics the Websocket connection
  
## Install

	go get -u github.com/abhinavxd/twitch-live-chat-downloader
  
## Getting started 

```go
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
			// Reconnect websocket connection
			err := twitch_chat.InitializeConnection(CHANNEL_NAME)
			if err != nil {
				log.Println(err)
			}
			continue
		}

		// Optionally you can Parse IRC tags which returns a `Message`
		// Or you can use the IRC tags directly
		parsedMsg, err := twitch_chat.ParseTags(message)
		
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("%+v\n\n", parsedMsg)
	}
}
```

![Screenshot from 2022-01-01 15-44-34](https://user-images.githubusercontent.com/48166553/147848396-b6c40ce9-87bf-42d0-8a41-ee1c9949c902.png)


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
