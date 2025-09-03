package main

import "net/url"

func normalizeURL(inURL string) (string, error) {
	outURL, err := url.Parse(inURL)
	return outURL.Hostname() + outURL.Path, err
}
