package diary

import (
	"strings"
	"time"
)

type Diary struct {
	// Relations
	Relations Relations
	Pages     []Page
}

type Relations map[string]Entry

func (relations Relations) Resolve(entry Entry) Entry {
	parts := entry.Parts()
	for i := len(parts) - 1; i >= 0; i-- {
		if related, ok := relations[parts[i]]; ok {
			relatedEntryArray := append(related.Parts(), parts[i+1:]...)
			relatedEntryString := strings.Join(relatedEntryArray, " > ")
			return Entry(relatedEntryString)
		}
	}
	return entry
}

type Page struct {
	Date       time.Time
	Title      string
	Categories Categories
}

type Categories map[string]Entries

type Entries []Entry

func (a Entries) Len() int           { return len(a) }
func (a Entries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Entries) Less(i, j int) bool { return a[i] < a[j] }

type Entry string

func (entry Entry) Parts() []string {
	parts := strings.Split(string(entry), ">")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

func (entry Entry) Prefixes() []string {
	parts := entry.Parts()
	prefixes := []string{}
	for i := 0; i < len(parts); i++ {
		prefix := strings.Join(parts[:i+1], " > ")
		prefixes = append(prefixes, prefix)
	}

	return prefixes
}
