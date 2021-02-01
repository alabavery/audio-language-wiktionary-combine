package definition

import (
	"audio-language/wiktionary/combine/util"
	"fmt"
)

type item struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Definitions  []string `json:"definitions"`
}

// DefinitionsWrapper wraps the content of a definitions file
type DefinitionsWrapper struct {
	Word       string
	Content    []item
	HasContent bool
}

// GetDefinitions instantiates and returns a DefinitionsWrapper pointer
func GetDefinitions(wordList []string, definitionsDir string) map[string]*DefinitionsWrapper {
	m := map[string]*DefinitionsWrapper{}
	for _, word := range wordList {
		def := DefinitionsWrapper{
			Word:       word,
			HasContent: false,
		}
		content := []item{}
		gotContent := util.GetJSONWhenFileMayNotExist(fmt.Sprintf("%v/%v.json", definitionsDir, word), &content)
		def.Content = content
		def.HasContent = gotContent
		m[word] = &def
	}
	return m
}
