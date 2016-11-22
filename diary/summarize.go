package diary

// SummarizeEntries summarizes the entries of the whole diary
func SummarizeEntries(diary *Diary) map[Entry]int {
	summary := map[Entry]int{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				entry := diary.Relations.Resolve(entry)
				summary[entry]++
			}
		}
	}
	return summary
}

// SummarizeCategoryEntries summarizes the parts of only the specified category.
func SummarizeCategoryEntries(diary *Diary, category string) map[Entry]int {
	summary := map[Entry]int{}
	for _, page := range diary.Pages {
		entries := page.Categories[category]
		for _, entry := range entries {
			entry := diary.Relations.Resolve(entry)
			summary[entry]++
		}
	}
	return summary
}

// SummarizeParts summarizes the parts of the whole diary
func SummarizeParts(diary *Diary) map[Entry]int {
	summary := map[Entry]int{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				entry = diary.Relations.Resolve(entry)
				for _, part := range entry.Prefixes() {
					summary[Entry(part)]++
				}
			}
		}
	}
	return summary
}

// SummarizeCategoryParts summarizes the parts of only the specified category.
func SummarizeCategoryParts(diary *Diary, category string) map[Entry]int {
	summary := map[Entry]int{}
	for _, page := range diary.Pages {
		entries := page.Categories[category]
		for _, entry := range entries {
			entry = diary.Relations.Resolve(entry)
			for _, part := range entry.Prefixes() {
				summary[Entry(part)]++
			}

		}
	}
	return summary
}

// SummarizeDiary summarizes and prints the whole diary
func SummarizeDiary(journal *Diary) {
	PrintSummarized(SummarizeEntries(journal))
}
