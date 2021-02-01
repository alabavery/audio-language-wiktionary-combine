package getflags

import (
	"flag"
	"fmt"
)

// FlagValues are the variables file paths necessary for the program
type FlagValues struct {
	Words       string
	Lemmas      string
	Definitions string
	Target      string
}

// GetFlags gets command line flags
func GetFlags() *FlagValues {
	wordListFilePathPtr := flag.String("words", "", "the path to the file containing the targeted words")
	lemmasDirPathPtr := flag.String("lemmas", "", "the path to the directory containing the lemmas")
	definitionsDirPtr := flag.String("definitions", "", "the path to the directory containing the definitions")
	targetDirPtr := flag.String("target", "", "the path to the target directory")
	flag.Parse()

	wordListFilePath := *wordListFilePathPtr
	lemmasDirPath := *lemmasDirPathPtr
	definitionsDir := *definitionsDirPtr
	targetDir := *targetDirPtr

	if wordListFilePath == "" || lemmasDirPath == "" || definitionsDir == "" || targetDir == "" {
		fmt.Println("Must provide the following flags:")
		flag.PrintDefaults()
		panic("missing flags")
	}
	return &FlagValues{
		Words:       wordListFilePath,
		Lemmas:      lemmasDirPath,
		Definitions: definitionsDir,
		Target:      targetDir,
	}
}
