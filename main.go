package main

import (
	"flag"
	"fmt"

	"github.com/mu-wahba/fily/utils"
)

var (
	filepath   *string
	word       *string
	ln         *bool
	searchType *string
	distpath   *string
	chuncksize *int64
	oldword    *string
	newword    *string
	k          *int
)

func init() {

	// Define flags in the init function
	searchType = flag.String("t", "", "type should be search, split")
	filepath = flag.String("fp", "", "FilePath to the file")
	distpath = flag.String("dp", "", "distination path to the file")
	word = flag.String("w", "", "Word to search for")
	chuncksize = flag.Int64("cs", 1, "chuncksize in MB")
	oldword = flag.String("ow", "", "word to replace")
	newword = flag.String("nw", "", "new word to replace ow")
	ln = flag.Bool("ln", false, "Line number flag")
	k = flag.Int("k", 1, "Top k ")

	flag.Parse() // Parse the flags

}

func main() {
	//search function
	switch *searchType {
	case "search":
		result, _ := utils.Search(*filepath, *word, *ln)
		if len(result) > 0 {
			for _, res := range result {
				fmt.Println(res)
			}
		} else {
			fmt.Println("Not Found")
		}
	case "split":
		utils.Split(*filepath, *distpath, *chuncksize)
	case "count":
		utils.Counter(*filepath)
	case "replace":
		utils.Replace(*filepath, *oldword, *newword)
	case "topwords":
		utils.TopKWords(*filepath, *k)
	default:
		fmt.Println("Please add -t to search, split, count, replace, topwords")
	}
}
