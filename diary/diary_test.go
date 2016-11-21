package diary_test

import (
	"testing"
	"time"


)

var example = &diary.Diary{
	Relations: diary.Relations{
		"John time":   "SOCIAL > friends > John time",
		"Jane time":   "SOCIAL > friends > Jane time",
		"Python time": "PRO > Python time",
	},
	Pages: []*diary.Page{
		&diary.Page{
			Date: time.Date(2016, 11, 19, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"joy": diary.Entries{
					"SOCIAL > friends > John time",
					"SOCIAL > friends > Nele time",
				},
			},
		},
		&diary.Page{
			Date: time.Date(2016, 10, 16, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"joy": diary.Entries{
					"SOCIAL > friends > Liina time",
					"SOCIAL > friends > Jane time",
					"weather nice",
					"Python time",
				},
			},
		},
		&diary.Page{
			Date: time.Date(2016, 10, 17, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"joy": diary.Entries{
					"Nele time",
					"SOCIAL > friends > John time",
					"lunch tasty",
					"Python time",
				},
			},
		},
	},
}

func TestSummarizeEntries(t *testing.T) {
	summary := diary.SummarizeEntries(example)

	expected := []struct {
		Entry diary.Entry
		Count int
	}{
		{"SOCIAL > friends > Nele time", 1},
		{"PRO > Python time", 2},
		{"lunch tasty", 1},
	}

	for i, test := range expected {
		if summary[test.Entry] != test.Count {
			t.Errorf("%d. %v %v != %v", i, test.Entry, summary[test.Entry], test.Count)
		}
	}
}

func TestSummarizeParts(t *testing.T) {
	summary := diary.SummarizeParts(example)
	expected := []struct {
		Entry diary.Entry
		Count int
	}{
		{"SOCIAL", 5},
		{"PRO", 2},
	}

	for i, test := range expected {
		if summary[test.Entry] != test.Count {
			t.Errorf("%d. %v %v != %v", i, test.Entry, summary[test.Entry], test.Count)
		}
	}
}
