package main

import (
	"fmt"
	"strings"
)

func receive(cfg *config) {
	<-cfg.concurrencyControl
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	if cfg.checkMaxPages() {
		return
	}
	defer receive(cfg)
	cfg.concurrencyControl <- struct{}{}
	if !isSameDomain(cfg.baseURL, rawCurrentURL) {
		return
	}
	curURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("\033[31m%v\033[37m\n", err)
		return
	}
	cfg.mu.Lock()
	if _, ok := cfg.pages[curURL]; ok {
		cfg.pages[curURL]++
		cfg.mu.Unlock()
		return
	} else {
		cfg.pages[curURL] = 1
	}
	cfg.pages["total"]++
	cfg.mu.Unlock()
	html, err := getHTML(curURL)
	if err != nil {
		fmt.Printf("\033[31m%v\033[37m\n", err)
		return
	}
	fmt.Printf("\033[32mGot html for URL: %s\033[37m\n", curURL)
	newURLs, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("\033[31mError getting urls from html: %v\033[37m\n", err)
		return
	}
	for _, url := range newURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}

func isSameDomain(baseURL, curURL string) bool {
	if !strings.Contains(curURL, "http") {
		return true
	}
	return strings.Contains(curURL, baseURL)
}

func (cfg *config) checkMaxPages() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return cfg.pages["total"] >= cfg.maxPages
}
