package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("Error making http request: %v", err)
	}
	req.Header.Add("User-Agent", "BootCrawler/1.0")
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error getting http response: %v", err)
	}
	ctype := res.Header.Get("content-type")
	if !strings.Contains(ctype, "text/html") {
		return "", fmt.Errorf("Error in http response, invalid content-type header : %s", ctype)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading http response: %v", err)
	}

	return string(data), nil
}
