package main

import (
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	// htmlReader := strings.NewReader(htmlBody)
	// node, err := html.Parse(htmlReader)
	// if err != nil {
	// return urls, err
	// }
	htmlLines := strings.SplitSeq(htmlBody, "\n")
	for line := range htmlLines {
		if strings.Contains(line, "href=") {
			url := strings.Split(line, "\"")[1]
			if !strings.Contains(url, "http") {
				url = rawBaseURL + url
			}
			urls = append(urls, url)
		}
	}

	return urls, nil
}
