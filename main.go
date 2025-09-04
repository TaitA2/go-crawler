package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		log.Fatalf("no arguments provided")
	}
	if len(args) > 3 {
		log.Fatalf("too many arguments provided")
	}
	rawURL, maxConcurrency, maxPages := getParams(args)
	fmt.Println("starting crawl of:", rawURL)

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            rawURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}
	cfg.pages["total"] = 0
	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()
	printReport(cfg.pages, cfg.baseURL)
}

func getParams(args []string) (url string, maxConcurrency, maxPages int) {
	url = args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Invalid max concurrency: %v", err)
	}
	maxPages, err = strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("Invalid max pages: %v", err)
	}
	return
}
