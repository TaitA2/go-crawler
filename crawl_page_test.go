package main

import "testing"

func TestIsSameDomain(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		curURL   string
		expected bool
	}{
		{
			name:     "same",
			baseURL:  "http://test.com",
			curURL:   "http://test.com",
			expected: true,
		}, {
			name:     "relative url",
			baseURL:  "http://test.com",
			curURL:   "path/one",
			expected: true,
		}, {
			name:     "extended url",
			baseURL:  "http://test.com",
			curURL:   "http://test.com/path/two",
			expected: true,
		}, {
			name:     "different url",
			baseURL:  "http://test.com",
			curURL:   "http://different.com/path/two",
			expected: false,
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSameDomain(tc.baseURL, tc.curURL)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}

		})
	}

}
