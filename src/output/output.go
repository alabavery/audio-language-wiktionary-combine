package output

import (
	"audio-language/wiktionary/combine/definition"
	"audio-language/wiktionary/combine/lemma"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type lemmaItem struct {
	Word        string   `json:"word"`
	Definitions []string `json:"definitions"`
}

type contentItem struct {
	PartOfSpeech string      `json:"part_of_speech"`
	Lemmas       []lemmaItem `json:"lemmas"`
}

// Wrapper has the final content
type Wrapper struct {
	Word       string
	Content    []contentItem
	HasContent bool
}

// GetOutputWrapper provides an OutputWrapper
func GetOutputWrapper(
	lemmaData *lemma.LemmasWrapper,
	definitions map[string]*definition.DefinitionsWrapper,
) *Wrapper {
	wrapper := &Wrapper{
		Word:       lemmaData.Word,
		HasContent: false,
		Content:    []contentItem{},
	}
	// the lemmaData should only not have content if the initial parse of the page was unsuccessful, or
	// if the initial page download was unsuccessful
	if lemmaData.HasContent {
		for _, partOfSpeechData := range lemmaData.Content {
			ci := &contentItem{
				PartOfSpeech: partOfSpeechData.PartOfSpeech,
				Lemmas:       []lemmaItem{},
			}
			// look up the definition(s) for each lemma of the word
			for _, lemma := range partOfSpeechData.Lemmas {
				li := lemmaItem{
					Word:        lemma,
					Definitions: []string{},
				}
				def, haveDefinition := definitions[lemma]
				if haveDefinition && def.HasContent {
					for _, item := range def.Content {
						if item.PartOfSpeech == partOfSpeechData.PartOfSpeech {
							li.Definitions = append(li.Definitions, item.Definitions...)
						}
					}
				}
				ci.Lemmas = append(ci.Lemmas, li)
			}
			wrapper.Content = append(wrapper.Content, *ci)
			wrapper.HasContent = true
		}
	}
	return wrapper
}

// Save the Content of the output
func (o *Wrapper) Save(targetDirectory string) {
	if o.HasContent {
		out, err := json.Marshal(o.Content)
		if err != nil {
			panic("Could not marshal json")
		}
		err = ioutil.WriteFile(fmt.Sprintf("%v/%v.json", targetDirectory, o.Word), out, 0644)
		if err != nil {
			panic("could not save file")
		}
	}
}
