package main

import (
	"fmt"
	"time"

	"github.com/tkivisik/gratitudelog/diary"
)

var example = diary.Diary{
	Relations: diary.Relations{
		"malle aeg":  "sots > sobrad > malle aeg",
		"peeter aeg": "sots > sobrad > peeter aeg",
		"Python aeg": "sots > Python aeg",
	},
	Pages: []diary.Page{
		diary.Page{
			Date: time.Date(2016, 11, 19, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"rõõm": diary.Entries{
					"sots > sobrad > marko aeg",
					"sots > sobrad > peeter aeg",
				},
			},
		},
		diary.Page{
			Date: time.Date(2016, 10, 16, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"rõõm": diary.Entries{
					"sots > sobrad > malle aeg",
					"sots > sobrad > peeter aeg",
					"ilm ilus",
					"Python aeg",
				},
			},
		},
		diary.Page{
			Date: time.Date(2016, 10, 17, 17, 46, 0, 0, time.UTC),
			Categories: diary.Categories{
				"rõõm": diary.Entries{
					"malle aeg",
					"sots > sobrad > peeter aeg",
					"ilm tuuline",
					"Python aeg",
				},
			},
		},
	},
}

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Print("ENTRIES: ")
	fmt.Println(diary.SummarizeEntries(&example))
	fmt.Println()
	fmt.Print("ENTRIES: ")
	fmt.Println(diary.SummarizeParts(&example))
}
