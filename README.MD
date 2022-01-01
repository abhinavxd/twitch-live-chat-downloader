# twitch-live-chat-downloader

Fetch twitch chat feed without any authentication

## How does it work?
* Mimics the Websocket connection in browser
  
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
```

![Screenshot from 2022-01-01 15-35-02](https://user-images.githubusercontent.com/48166553/147848280-7a5daaf1-a03c-4e3d-a93f-f4817cddcda9.png)


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
