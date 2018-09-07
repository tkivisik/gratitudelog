package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkivisik/gratitudelog/diary"
)

func printHelp() {
	fmt.Println("Enjoy the feeling of gratitude.\n")
	fmt.Println("Argument options:")
	fmt.Println(" * summarize or summarizeGratitude")
	/*
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

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "summarize":
			diary.SummarizeDiary(journal)
		case "summarizeGratitude":
			fmt.Println(diary.SummarizeCategoryParts(journal, defaultCategory))
		default:
			fmt.Println("Default: ")
			printHelp()
		}
	}
}
