package diary

import (
	"fmt"
	"strings"
)

/*
It's meaningful to summarize:
* All entries (with relations) in a diary SummarizeEntries
* All parts (including SOCIAL) in a diary SummarizeParts
* All entries in a certain category SummarizeCategoryEntries
* All parts in a certain category SummarizeCategoryParts
*/

const DotWeight = 0.5

// SummarizeEntries summarizes the entries of the whole diary
func SummarizeEntries(diary *Diary) map[Entry]float32 {
	summary := map[Entry]float32{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				entry := diary.Relations.Resolve(entry)
				if strings.HasPrefix(string(entry), ".") {
					summary[entry] = summary[entry] + DotWeight
				} else {
					summary[entry]++
				}
			}
		}
	}
	return summary
}

// SummarizeCategoryEntries summarizes the parts of only the specified category.
func SummarizeCategoryEntries(diary *Diary, category string) map[Entry]float32 {
	summary := map[Entry]float32{}
	for _, page := range diary.Pages {
		entries := page.Categories[category]
		for _, entry := range entries {
			entry := diary.Relations.Resolve(entry)
			if strings.HasPrefix(string(entry), ".") {
				summary[entry] = summary[entry] + DotWeight
			} else {
				summary[entry]++
			}
		}
	}
	return summary
}

// SummarizeParts summarizes the parts of the whole diary
func SummarizeParts(diary *Diary) map[Entry]float32 {
	summary := map[Entry]float32{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				entry = diary.Relations.Resolve(entry)
				for _, part := range entry.Prefixes() {
					if strings.HasPrefix(part, ".") {
						summary[Entry(part[1:])] += DotWeight
					} else {
						summary[Entry(part)]++
					}
				}
			}
		}
	}
	return summary
}

// SummarizeCategoryParts summarizes the parts of only the specified category.
func SummarizeCategoryParts(diary *Diary, category string) map[Entry]float32 {
	summary := map[Entry]float32{}
	for _, page := range diary.Pages {
		entries := page.Categories[category]
		for _, entry := range entries {
			entry = diary.Relations.Resolve(entry)
			for _, part := range entry.Prefixes() {
				if strings.HasPrefix(part, ".") {
					summary[Entry(part[1:])] = summary[Entry(part[1:])] + DotWeight
				} else {
					summary[Entry(part)]++
				}
			}
		}
	}
	return summary
}

// SummarizeDiary summarizes and prints the whole diary
func SummarizeDiary(journal *Diary) {
	fmt.Println("Category Entries")
	PrintSummarized(SummarizeCategoryEntries(journal, "gratitude"))
	fmt.Println()
	fmt.Println("Category Parts")
	PrintSummarized(SummarizeCategoryParts(journal, "gratitude"))

}
