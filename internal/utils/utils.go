package utils

import (
	"bufio"
	"errors"
	"math/rand/v2"
	"os"
	"slices"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func GetRandomWord(length int) (string, error) {
	if length < 4 || length > 11 {
		return "", errors.New("Length must be between 5 and 10")
	}

	words, err := ReadLines("internal/data/words")
	if err != nil {
		return "", err
	}

	filtered := slices.Collect(func(yield func(string) bool) {
		for _, word := range words {
			if len(word) == length {
				if !yield(word) {
					return
				}
			}
		}
	})

	return filtered[rand.IntN(len(filtered))], nil
}
