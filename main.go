package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	file, err := os.Open("./words.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	wordCount := CountWords(file)
	fmt.Println("File contains:", wordCount, "words.")
}

func CountWords(file io.Reader) int {
	wordsCount := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordsCount++
	}

	return wordsCount
}
