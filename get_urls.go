package main

import (
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	htmlLines := strings.SplitSeq(htmlBody, ">")
	for line := range htmlLines {
		if strings.Contains(line, "href=\"") {
			url := strings.Split(line, "\"")[1]
			if !strings.Contains(url, "http") {
				url = rawBaseURL + url
			}
			urls = append(urls, url)
		}
	}

	return urls, nil
}
