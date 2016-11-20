package diary

import (
	"strings"
	"time"
)

type Diary struct {
	// Relations
	Pages []Page
}

type Page struct {
	Date       time.Time
	Title      string
	Categories Categories
}

type Categories map[string]Entries

type Entries []Entry

type Entry string

func (entry Entry) Parts() []string {
	parts := strings.Split(string(entry), ">")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
