package visualisation

// WordCloud makes a wordcloud of the whole diary
import (
	"bufio"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/marcusolsson/wordcloud"
	"github.com/tkivisik/gratitudelog/diary"
	"golang.org/x/image/font/gofont/goregular"
)

// WordCloud creates a ~wordcloud.png from key-value pairs, by category. Use 'all' to include all categories.
func WordCloud(journal *diary.Diary, category string) {
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
	if category != "all" {
		for entry, count := range diary.SummarizeCategoryParts(journal, category) {
			words[string(entry)] = count
		}
	} else {
		for entry, count := range diary.SummarizeParts(journal) {
			words[string(entry)] = count
		}
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
