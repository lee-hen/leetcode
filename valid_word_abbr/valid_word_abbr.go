package valid_word_abbr

import "fmt"

type ValidWordAbbr struct {
	dictionary map[string]bool
	wordAbbrCounter map[string]int
}

func Constructor(dic []string) ValidWordAbbr {
	dictionary := make(map[string]bool)
	wordAbbrCounter := make(map[string]int)
	for _, str := range dic {
		if !dictionary[str] {
			dictionary[str] = true
			key := generateKey(str)
			wordAbbrCounter[key]++
		}
	}

	return ValidWordAbbr {
		dictionary: dictionary,
		wordAbbrCounter: wordAbbrCounter,
	}
}

func (v *ValidWordAbbr) IsUnique(word string) bool {
	key := generateKey(word)
	if _, ok := v.dictionary[word]; !ok {
		return v.wordAbbrCounter[key] == 0
	}

	return v.wordAbbrCounter[key] == 1
}

func generateKey(str string) string {
	return fmt.Sprintf("%s:%d:%s", string(str[0]), len(str)-2, string(str[len(str)-1]))
}
