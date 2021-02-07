package lemma

import (
	"audio-language/wiktionary/combine/util"

	"github.com/ninetypercentlanguage/word-utils/lemma"
)

// LemmasWrapper wraps lemmas
type LemmasWrapper struct {
	Word       string
	Content    lemma.Content
	HasContent bool
}

// NewLemmasWrapper provides a wrapper around a word wrt its lemmas
func NewLemmasWrapper(word string, lemmasDir string) *LemmasWrapper {
	content, hasContent := getFileContent(word, lemmasDir)
	return &LemmasWrapper{
		Word:       word,
		HasContent: hasContent,
		Content:    content,
	}
}

// GetFlatLemmaList translates a bunch of wrappers into a flat list
func GetFlatLemmaList(wrappers []*LemmasWrapper) *[]string {
	var lemmas []string
	for _, wrapper := range wrappers {
		for _, content := range wrapper.Content {
			lemmas = append(lemmas, content.Lemmas...)
		}
	}
	return &lemmas
}

func getFileContent(word string, lemmasDir string) (lemma.Content, bool) {
	var content lemma.Content
	exists := util.GetJSONWhenFileMayNotExist(
		lemmasDir+"/"+word+".json",
		&content,
	)
	return content, exists
}
