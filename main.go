package main

import (
	"fmt"
	"os"
)

func countWords(data []byte) int {
	wordCount := 0
	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}
	wordCount++
	return wordCount
}

func main() {
	data, _ := os.ReadFile("./words.txt")
	fmt.Println("File contains:", countWords(data), "words.")
}
