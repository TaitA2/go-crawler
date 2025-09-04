package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("no website provided")
	}
	if len(args) > 1 {
		log.Fatalf("too many arguments provided")
	}
	rawURL := args[0]
	fmt.Println("starting crawl of:", rawURL)
	html, err := getHTML(rawURL)
	if err != nil {
		log.Fatalf("Error getting html: %v", err)
	}
	fmt.Println(html)
}
