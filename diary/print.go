package diary

import (
	"fmt"
	"sort"
)

func PrintSummarized(input map[Entry]int) {
	keys := make(Entries, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	for _, entry := range keys {
		fmt.Println(input[entry], " : ", entry)
	}
}
