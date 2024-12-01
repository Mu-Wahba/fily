package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Counter(fp string) error {
	//openfile
	file, err := os.Open(fp)
	check(err)
	defer file.Close()
	wordsCount, linesCounet := int64(0), int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesCounet++
		words := strings.Fields(line)
		wordsCount = wordsCount + int64(len(words))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wc := fmt.Sprintf("Words Count: %d", wordsCount)
	lc := fmt.Sprintf("Lines Count: %d", linesCounet)
	fmt.Println(wc)
	fmt.Println(lc)
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal("Couldn't open file ", err.Error())
	}
}
