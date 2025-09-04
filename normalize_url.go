package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inURL string) (string, error) {
	outURL, err := url.Parse(inURL)
	return strings.ReplaceAll(outURL.Hostname()+outURL.Path, "//", "/"), err
}
