package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Search(filepath string, word string, ln bool) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("couldn't open file", err)
		os.Exit(1)
	}
	defer file.Close()
	result := []string{}
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if strings.Contains(line, word) {
			if ln {
				result = append(result, fmt.Sprintf("%d >> %s", lineCount, line))
			} else {
				result = append(result, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return result, nil
}
