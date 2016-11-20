package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkivisik/gratitudelog/diary"
)

func printHelp() {
	fmt.Println("Usage: ...")
	os.Exit(1)

}

func main() {
	if len(os.Args) < 3 {
		printHelp()
	}
	data, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	journal, err := diary.ParseJSON(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "summarize":
		SummarizeDiary(journal)
	default:
		printHelp()
	}
}

// SummarizeDiary summarizes the whole diary
func SummarizeDiary(journal *diary.Diary) {
	diary.PrintSummarized(diary.SummarizeEntries(journal))

}
