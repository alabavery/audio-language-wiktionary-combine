package output

import (
	"audio-language/wiktionary/combine/definition"
	"audio-language/wiktionary/combine/lemma"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ninetypercentlanguage/word-utils/combined"
)

// Wrapper around the final output
type Wrapper struct {
	Content    combined.Content
	HasContent bool
	Word       string
}

// GetOutputWrapper provides an OutputWrapper
func GetOutputWrapper(
	lemmaData *lemma.LemmasWrapper,
	definitions map[string]*definition.DefinitionsWrapper,
) *Wrapper {
	wrapper := &Wrapper{
		Word:       lemmaData.Word,
		HasContent: false,
		Content:    combined.Content{},
	}
	// the lemmaData should only not have content if the initial parse of the page was unsuccessful, or
	// if the initial page download was unsuccessful
	if lemmaData.HasContent {
		for _, partOfSpeechData := range lemmaData.Content {
			ci := &combined.ContentItem{
				PartOfSpeech: partOfSpeechData.PartOfSpeech,
				Lemmas:       []combined.LemmaItem{},
			}
			// look up the definition(s) for each lemma of the word
			for _, lemma := range partOfSpeechData.Lemmas {
				li := combined.LemmaItem{
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

// Save the the output
func (o *Wrapper) Save(targetDirectory string, dryRun bool) {
	if o.HasContent {
		var content combined.Content = o.Content
		out, err := json.Marshal(content)
		if err != nil {
			panic("Could not marshal json")
		}
		if dryRun {
			fmt.Printf("output: %v\n", o)
			return
		}
		err = ioutil.WriteFile(fmt.Sprintf("%v/%v.json", targetDirectory, o.Word), out, 0644)
		if err != nil {
			panic("could not save file")
		}
	}
}
