package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var words = map[string]int{}

// TODO Implement Heap to get Top K words instead on sorting
func TopKWords(fp string, k int) error {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal("Couldn't open file", err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}
	res := sortByValue(words)
	for i := 0; i < k; i++ {
		fmt.Printf("%s, %d\n", res[i], words[res[i]])
	}
	return nil
}

func sortByValue(kvmap map[string]int) []string {
	keys := make([]string, 0, len(kvmap))
	for key := range kvmap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return kvmap[keys[i]] > kvmap[keys[j]]
	})

	return keys

}
