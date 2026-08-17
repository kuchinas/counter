package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	totalWords := 0
	filenames := os.Args[1:]
	didError := false
	for _, filename := range filenames {
		wordCount, err := CountWordsInFile(filename)
		if err != nil {
			didError = true
			_, _ = fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}
		totalWords += wordCount
		fmt.Println(wordCount, filename)
	}
	if len(filenames) == 0 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount)
	}
	if len(filenames) > 1 {
		fmt.Println(totalWords, "total")
	}
	if didError {
		os.Exit(1)
	}
}
