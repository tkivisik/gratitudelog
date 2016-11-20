package diary

func SummarizeEntries(diary *Diary) map[Entry]int {
	summary := map[Entry]int{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				summary[entry]++
			}
		}
	}
	return summary
}

func SummarizeParts(diary *Diary) map[string]int {
	summary := map[string]int{}
	for _, page := range diary.Pages {
		for _, entries := range page.Categories {
			for _, entry := range entries {
				for _, part := range entry.Parts() {
					summary[part]++
				}
			}
		}
	}
	return summary
}
