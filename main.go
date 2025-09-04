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
	pages := make(map[string]int)
	crawlPage(rawURL, rawURL, pages)
	for k, v := range pages {
		fmt.Println(k, " - ", v)
	}
}
