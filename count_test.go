package main_test

import (
	"strings"
	"testing"

	counter "github.com/kuchinas/counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{name: "five words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "Empty inputs",
			input: "",
			wants: 0,
		},
		{
			name:  "A space character",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one, two, three\nfour, five",
			wants: 5,
		},
		{
			name:  "multiply spaces",
			input: "This is a sentence.  This is another",
			wants: 7,
		},
		{
			name:  "Prefixed multiply spaces",
			input: "  Hello",
			wants: 1,
		},
		{
			name:  "Suffixzed multiply spaces",
			input: "Hello  ",
			wants: 1,
		},
		{
			name:  "Tab character",
			input: "Hello\tWorld",
			wants: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result := counter.CountWords(reader)
			if result != tc.wants {
				t.Errorf("countWords(%v) = %v, wants %v", tc.input, result, tc.wants)
			}
		})
	}
}
