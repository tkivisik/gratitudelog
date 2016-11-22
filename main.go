package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkivisik/gratitudelog/diary"
	"github.com/tkivisik/gratitudelog/visualisation"
)

func printHelp() {
	fmt.Println("Usage: XXXX")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
	}

	// Read in a JSON file
	data, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		os.Exit(1)
	}

	// Process that JSON file
	journal, err := diary.ParseJSON(data)
	if err != nil {
		fmt.Println("Here")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "summarize":
		diary.SummarizeDiary(journal)
	case "summarizeJoy":
		fmt.Println(diary.SummarizeCategoryParts(journal, "rõõm"))
	case "wordcloud":
		visualisation.WordCloud(journal)
	default:
		fmt.Println("Default: ")
		printHelp()
	}
}
