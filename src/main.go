package main

import (
	"audio-language/wiktionary/combine/definition"
	"audio-language/wiktionary/combine/getflags"
	"audio-language/wiktionary/combine/lemma"
	"audio-language/wiktionary/combine/output"
	"audio-language/wiktionary/combine/word"
)

func main() {
	flagVals := getflags.GetFlags()
	// get list of words from original word list file
	wordList := word.GetWords(flagVals.Words)
	var lemmas []*lemma.LemmasWrapper
	for _, word := range wordList {
		lemmaData := lemma.NewLemmasWrapper(word, flagVals.Lemmas)
		lemmas = append(lemmas, lemmaData)
	}
	flatLemmas := lemma.GetFlatLemmaList(lemmas)
	definitions := definition.GetDefinitions(*flatLemmas, flagVals.Definitions)

	for _, word := range wordList {
		// inefficient to go and get lemmas again, but we are in no rush
		lemmaData := lemma.NewLemmasWrapper(word, flagVals.Lemmas)
		o := output.GetOutputWrapper(lemmaData, definitions)
		output.Save(o, flagVals.Target, flagVals.DryRun)
	}
}
