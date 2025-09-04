package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\nREPORT for %s\n============================\n", baseURL)
	keys := sortKeys(pages)
	for _, k := range keys {
		v := pages[k]
		if k != "total" {
			fmt.Printf("Found %d internal links to %s\n", v, k)
		}
	}
}

func sortKeys(pages map[string]int) (sortedKeys []string) {
	keys := make([]string, 0, len(pages))

	for key := range pages {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return pages[keys[i]] < pages[keys[j]]
	})
	return keys
}
