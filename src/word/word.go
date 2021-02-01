package word

import (
	"audio-language/wiktionary/combine/util"
)

// GetWords reads the words from the words file
func GetWords(filePath string) []string {
	var content []string
	exists := util.GetJSONWhenFileMayNotExist(filePath, &content)
	if !exists {
		panic("could not get words file")
	}
	return content
}
