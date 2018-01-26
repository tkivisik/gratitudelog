package diary

import (
	"fmt"
	"sort"
)

// PrintSummarized prints formatted key-value pairs.
func PrintSummarized(input map[Entry]float32) {
	keys := make(Entries, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	for _, entry := range keys {
		fmt.Printf("%5.1f : %v\n", input[entry], entry)
	}
}
