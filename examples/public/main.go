package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/zekrotja/godeepl"
)

func main() {
	godotenv.Load()

	c := godeepl.New()

	res, err := c.Translate(
		godeepl.LangAuto, godeepl.LangEnglish,
		"Jo, was geht ab du alter Sack! Dauert das noch lange?",
		godeepl.TranslationOptions{
			Formality:        godeepl.FormalityFormal,
			PreferedNumBeams: 3,
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Detected Target Lang:", res.SourceLang)

	fmt.Print("\nTranslated sentences:\n\n")
	for _, trans := range res.Translations {
		fmt.Println(trans.Beam(0).ProcessedSentence)
		if len(trans.Beams) > 1 {
			for _, beam := range trans.Beams[1:] {
				fmt.Println("  -", beam.ProcessedSentence)
			}
		}
	}

	fmt.Println("\nPreferred Result:\n", res.Assemble())
}
