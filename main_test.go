package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five\nsix seven eight"
	wants := 8
	result := countWords([]byte(input))
	if result != wants {
		t.Errorf("countWords(%q) = %d, want %d", input, result, wants)
	}
}
