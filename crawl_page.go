package main

import (
	"fmt"
	"log"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !isSameDomain(rawBaseURL, rawCurrentURL) {
		return
	}
	curURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("\033[31m%v\033[37m", err)
		return
	}
	if _, ok := pages[curURL]; ok {
		pages[curURL]++
		return
	} else {
		pages[curURL] = 1
	}
	html, err := getHTML(curURL)
	if err != nil {
		log.Printf("\033[31m%v\033[37m", err)
		return
	}
	fmt.Println("Got html for URL: ", curURL)
	newURLs, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		log.Printf("Error getting urls from html: %v", err)
		return
	}
	for _, url := range newURLs {
		crawlPage(rawBaseURL, url, pages)
	}

}

func isSameDomain(baseURL, curURL string) bool {
	if !strings.Contains(curURL, "http") {
		return true
	}
	return strings.Contains(curURL, baseURL)
}
