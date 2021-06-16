package godeepl

import "strings"

func (tr *TranslationResult) Sentences() []string {
	sentences := make([]string, 0, len(tr.Translations))

	for _, t := range tr.Translations {
		if len(t.Beams) > 0 {
			sentences = append(sentences, t.Beams[0].ProcessedSentence)
		}
	}

	return sentences
}

func (tr *TranslationResult) Assemble() string {
	return strings.Join(tr.Sentences(), " ")
}
