// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tkivisik/gratitudelog/diary"
)

var (
	input  = flag.String("input", "", "input file")
	output = flag.String("output", "", "output file")
)

func main() {
	flag.Parse()

	if *input == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *output == "" {
		*output = *input + ".json"
	}

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	rxPage := regexp.MustCompile(`^[ETKNRLPMWFS](\d\d\.\d\d.\d\d\d\d)\s*\-\s*(.*)$`)
	rxCategory := regexp.MustCompile(`^(\p{L}+)\s*:\s*(.*)$`)

	var journal diary.Diary
	var page *diary.Page
	var category string
	var allentries string

	flushentries := func() {
		entries := diary.Entries{}

		allentries = strings.TrimSpace(allentries)
		for _, entry := range strings.Split(allentries, ",") {
			entry = strings.TrimSpace(entry)
			entry = strings.TrimRight(entry, ".")
			entry = strings.TrimLeft(entry, "(")
			if entry == "" || entry == "-" {
				continue
			}
			entries = append(entries, diary.Entry(entry))
		}
		if len(entries) > 0 {
			page.Categories[category] = entries
		}
		allentries = ""
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimRight(line, "\n\r")
		if strings.TrimSpace(line) == "" {
			continue
		}

		if rxPage.MatchString(line) {
			flushentries()

			page = diary.NewPage()
			journal.Pages = append(journal.Pages, page)

			datetitle := rxPage.FindStringSubmatch(line)
			page.Date, err = time.Parse("02.01.2006", datetitle[1])
			if err != nil {
				log.Fatal(err)
			}
			page.Title = datetitle[2]

			category = ""
			continue
		}

		if rxCategory.MatchString(line) {
			flushentries()

			categoryentries := rxCategory.FindStringSubmatch(line)
			category = categoryentries[1]

			allentries = strings.TrimSpace(categoryentries[2])
			continue
		}

		allentries += " " + strings.TrimSpace(line)
	}
	flushentries()

	data, err = json.MarshalIndent(journal, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(*output, data, 0755)
}
