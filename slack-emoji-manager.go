package main

import (
	"fmt"
	"log"
	"os"
	"slack-emoji-manager/emoji"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Invalid arguments")
	}

	op := os.Args[1]
	token := os.Args[2]

	switch op {
	case "get":
		fmt.Println("Fetching emoji list")
		err := emoji.GetEmojis(token, false)
		if err != nil {
			panic(err)
		}
	case "download":
		fmt.Println("Downloading all emojis")
		err := emoji.GetEmojis(token, true)
		if err != nil {
			panic(err)
		}
	}
}
