package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	fmt.Println("File contains:", CountWords(data), "words.")
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	wordCount := 0
	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}
	wordCount++
	return wordCount
}
