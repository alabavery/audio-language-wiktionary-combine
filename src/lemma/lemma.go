package lemma

import "audio-language/wiktionary/combine/util"

type item struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Lemmas       []string `json:"lemmas"`
	Exists       bool     `json:"exists"`
}

// LemmasWrapper wraps lemmas
type LemmasWrapper struct {
	Word       string
	Content    []item
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

func getFileContent(word string, lemmasDir string) ([]item, bool) {
	var content []item
	exists := util.GetJSONWhenFileMayNotExist(
		lemmasDir+"/"+word+".json",
		&content,
	)
	return content, exists
}
