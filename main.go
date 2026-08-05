package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	fmt.Println("File contains:", CountWords(data), "words.")
}

func CountWords(data []byte) int {
	wordsCount := bytes.Fields(data)
	return len(wordsCount)
}
