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
	const maxConcurrency = 10
	cfg := config{
		pages:              make(map[string]int),
		baseURL:            rawURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}
	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()
	fmt.Println()
	for k, v := range cfg.pages {
		fmt.Println(k, " - ", v)
	}
}
