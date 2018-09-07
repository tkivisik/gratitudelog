package diary

import (
	"strings"
	"time"
)

// Diary contains Relations and Pages (days)
type Diary struct {
	Relations Relations
	Pages     []*Page
}

type Relations map[string]Entry

// Resolve returns an Entry with its full relation
func (relations Relations) Resolve(entry Entry) Entry {
	parts := entry.Parts()
	// go through the parts from the end to the beginning
	for i := len(parts) - 1; i >= 0; i-- {
		// until you find a part which is part of a relations system
		if related, ok := relations[parts[i]]; ok {
			// append everything from a relations system + parts from
			// this entry until that relation was found.
			relatedEntryArray := append(related.Parts(), parts[i+1:]...)
			relatedEntryString := strings.Join(relatedEntryArray, " > ")
			return Entry(relatedEntryString)
		}
	}
	return entry
}

// Page contains the date, title and categories for a day
type Page struct {
	Date        time.Time
	DateWritten time.Time
	Title       string
	Categories  Categories
}

func NewPage() *Page {
	return &Page{
		Date:        time.Time{},
		DateWritten: time.Time{},
		Title:       "",
		Categories:  make(Categories),
	}
}

// Categories contains all categories (e.g. gratitude, geo) and their content
type Categories map[string]Entries

// Entries contains a list of entries a category
type Entries []Entry

func (a Entries) Len() int           { return len(a) }
func (a Entries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Entries) Less(i, j int) bool { return a[i] < a[j] }

// Entry may or may not have a relation (e.g."lunch tasty", "SOCIAL > friends > John time")
type Entry string

// Parts returns parts of an Entry as a slice of strings
func (entry Entry) Parts() []string {
	parts := strings.Split(string(entry), ">")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

// Prefixes returns a string for every level of an entry.
func (entry Entry) Prefixes() []string {
	parts := entry.Parts()
	prefixes := []string{}
	for i := 0; i < len(parts); i++ {
		prefix := strings.Join(parts[:i+1], " > ")
		prefixes = append(prefixes, prefix)
	}
	return prefixes
}
