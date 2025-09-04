package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	cfg := config{
		pages:   make(map[string]int),
		baseURL: rawURL,
		mu:      &sync.Mutex{},
	}
	cfg.crawlPage(rawURL)
	for k, v := range cfg.pages {
		fmt.Println(k, " - ", v)
	}
}
