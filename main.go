package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font/gofont/goregular"

	"github.com/golang/freetype/truetype"
	"github.com/marcusolsson/wordcloud"
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
	case "summarizeJoy":
		fmt.Println(diary.SummarizeCategoryParts(journal, "rõõm"))
	case "wordcloud":
		WordCloud(journal)
	default:
		printHelp()
	}
}

// SummarizeDiary summarizes the whole diary
func SummarizeDiary(journal *diary.Diary) {
	diary.PrintSummarized(diary.SummarizeEntries(journal))

}

// WordCloud makes a wordcloud of the whole diary
func WordCloud(journal *diary.Diary) {
	defaultWidth := 1024
	defaultHeight := 1024
	defaultOutput := "~wordcloud.png"

	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the cloud.
	c := wordcloud.Cloud{
		Width:     defaultWidth,
		Height:    defaultHeight,
		Font:      f,
		Generator: wordcloud.NewSpiralGenerator(),
	}

	words := make(map[string]int)
	for entry, count := range diary.SummarizeCategoryParts(journal, "rõõm") {
		words[string(entry)] = count
	}

	im := c.Generate(words)

	// Save to file.
	if err := writeToFile(defaultOutput, im); err != nil {
		log.Fatal(err)
	}
}

func writeToFile(str string, im image.Image) error {
	f, err := os.Create(str)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	if err = png.Encode(w, im); err != nil {
		return err
	}

	if err = w.Flush(); err != nil {
		return err
	}

	return nil
}
