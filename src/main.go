package main

import (
	"audio-language/wiktionary/combine/definition"
	"audio-language/wiktionary/combine/getflags"
	"audio-language/wiktionary/combine/lemma"
	"audio-language/wiktionary/combine/output"
	"audio-language/wiktionary/combine/word"
)

type datapoint struct {
	name string
	data map[string]interface{}
}

type partOfSpeechData struct {
	name        string
	lemma       string
	definitions []string
	datapoints  []datapoint
}

type wordData struct {
	word          string
	partsOfSpeech []partOfSpeechData
}

func main() {
	flagVals := getflags.GetFlags()
	// get list of words from original word list file
	wordList := word.GetWords(flagVals.Words)
	definitions := definition.GetDefinitions(wordList, flagVals.Definitions)
	for _, word := range wordList {
		// get lemmas from the lemmas directory
		lemmaData := lemma.NewLemmasWrapper(word, flagVals.Lemmas)
		o := output.GetOutputWrapper(lemmaData, definitions)
		o.Save(flagVals.Target)
	}
}
