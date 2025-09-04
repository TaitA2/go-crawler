package main

import (
	"fmt"
	"log"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	if !isSameDomain(cfg.baseURL, rawCurrentURL) {
		return
	}
	curURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("\033[31m%v\033[37m", err)
		return
	}
	if _, ok := cfg.pages[curURL]; ok {
		cfg.pages[curURL]++
		return
	} else {
		cfg.pages[curURL] = 1
	}
	html, err := getHTML(curURL)
	if err != nil {
		log.Printf("\033[31m%v\033[37m", err)
		return
	}
	fmt.Println("Got html for URL: ", curURL)
	newURLs, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		log.Printf("Error getting urls from html: %v", err)
		return
	}
	for _, url := range newURLs {
		cfg.crawlPage(url)
	}

}

func isSameDomain(baseURL, curURL string) bool {
	if !strings.Contains(curURL, "http") {
		return true
	}
	return strings.Contains(curURL, baseURL)
}
