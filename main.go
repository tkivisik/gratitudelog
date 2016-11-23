package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkivisik/gratitudelog/diary"
	"github.com/tkivisik/gratitudelog/visualisation"
)

func printHelp() {
	fmt.Println("Enjoy the feeling of gratitude.")
	/* Menu not functional yet.
	fmt.Println("Click:")
	fmt.Println("'1' to enter the name of .json gratitude log file and read it in")
	fmt.Println("'2' to summarize the whole diary")
	fmt.Println("'3' to enter the category you want to summarize (e.g. 'gratitude')")
	fmt.Println("'4' to make a word cloud")
	*/
	//	os.Exit(1)
}

func main() {
	defaultFile := "sample-gratitude.json"
	defaultCategory := "gratitude"
	if len(os.Args) < 3 {
		printHelp()
	}

	// Read in a JSON file
	data, err := ioutil.ReadFile(defaultFile)
	if err != nil {
		os.Exit(1)
	}

	// Process that JSON file
	journal, err := diary.ParseJSON(data)
	if err != nil {
		os.Exit(1)
	}

	diary.SummarizeDiary(journal)

	switch os.Args[1] {
	case "summarize":
		diary.SummarizeDiary(journal)
	case "summarizeGratitude":
		fmt.Println(diary.SummarizeCategoryParts(journal, defaultCategory))
	case "wordcloud":
		visualisation.WordCloud(journal, defaultCategory)
	default:
		fmt.Println("Default: ")
		printHelp()
	}
}
