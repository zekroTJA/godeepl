package godeepl

import "strings"

// Sentences returns an array of the translated sentences
// using the first result beam of each sentence.
func (tr *TranslationResult) Sentences() []string {
	sentences := make([]string, 0, len(tr.Translations))

	for _, t := range tr.Translations {
		if len(t.Beams) > 0 {
			sentences = append(sentences, t.Beams[0].ProcessedSentence)
		}
	}

	return sentences
}

// Assemble takes all sentences of a translation using the first
// beam of each translation and assembles them to one string.
func (tr *TranslationResult) Assemble() string {
	return strings.Join(tr.Sentences(), " ")
}

// Translation is a safe getter for a translation with
// the given index. If there is no entry at the given
// index, nil is returned.
func (tr *TranslationResult) Translation(i int) *Translation {
	if tr == nil || i < 0 || len(tr.Translations) <= i {
		return nil
	}
	return tr.Translations[i]
}

// Beam is a safe getter for a beam with
// the given index. If there is no entry at the given
// index, nil is returned.
func (t *Translation) Beam(i int) *Beam {
	if t == nil || i < 0 || len(t.Beams) <= i {
		return nil
	}
	return t.Beams[i]
}
