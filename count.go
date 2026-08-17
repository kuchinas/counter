package main

import (
	"bufio"
	"io"
	"os"
)

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		return 0, err
	}
	return CountWords(file), nil
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
