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

	example2, err := diary.ParseJSON(`
{
    "Relations": {
        "Python aeg": "sots > Python aeg",
        "malle aeg": "sots > sobrad > malle aeg",
        "peeter aeg": "sots > sobrad > peeter aeg"
    },
    "Pages": [
        {
            "Date": "2016-11-19T17:46:00Z",
            "Title": "",
            "Categories": {
                "rõõm": [
                    "sots > sobrad > marko aeg",
                    "sots > sobrad > peeter aeg"
                ]
            }
        },
        {
            "Date": "2016-10-16T17:46:00Z",
            "Title": "",
            "Categories": {
                "rõõm": [
                    "sots > sobrad > malle aeg",
                    "sots > sobrad > peeter aeg",
                    "ilm ilus",
                    "Python aeg"
                ]
            }
        },
        {
            "Date": "2016-10-17T17:46:00Z",
            "Title": "",
            "Categories": {
                "rõõm": [
                    "malle aeg",
                    "sots > sobrad > peeter aeg",
                    "ilm tuuline",
                    "Python aeg"
                ]
            }
        }
    ]
}
`)

	fmt.Println()
	fmt.Println(err)
	fmt.Print("ENTRIES: ")
	fmt.Println(diary.SummarizeEntries(example2))
	fmt.Println()
	fmt.Print("PARTS: ")
	fmt.Println(diary.SummarizeParts(example2))
	fmt.Println()
	fmt.Println("SORTED PARTS: ")
	diary.PrintSummarized(diary.SummarizeParts(example2))
	fmt.Println()
	fmt.Println("SORTED ENTRIES: ")
	diary.PrintSummarized(diary.SummarizeEntries(example2))
}
