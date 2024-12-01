package utils

import (
	"fmt"
	"os"
	"regexp"
)

// Replace replaces all occurrences of a specific word in the file with a new word in the same file
func Replace(filePath string, oldWord string, newWord string) error {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	dataBytes, err := os.ReadFile(filePath)
	check(err)

	re := regexp.MustCompile(`\b` + regexp.QuoteMeta(oldWord) + `\b`)
	if !re.Match(dataBytes) {
		fmt.Printf("The word '%s' was not found in the file.\n", oldWord)
		return nil

	}
	data := re.ReplaceAllString(string(dataBytes), newWord)

	if err := os.WriteFile(filePath, []byte(data), os.ModePerm); err != nil {
		check(err)
	}
	fmt.Printf("Replaces %s with %s", oldWord, newWord)
	return nil
}
