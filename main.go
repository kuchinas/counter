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
	totalWords := 0
	if len(os.Args) < 2 {
		log.Fatal("no filename provided")
	}
	filenames := os.Args[1:]
	for _, filename := range filenames {
		wordCount := CountWordsInFile(filename)
		totalWords += wordCount
		fmt.Println(wordCount, filename)
	}
	if len(filenames) > 1 {
		fmt.Println(totalWords, "total")
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	return CountWords(file)
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
