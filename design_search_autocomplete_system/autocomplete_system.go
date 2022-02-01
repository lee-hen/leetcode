package design_search_autocomplete_system

import (
	"sort"
	"strings"
)

type autocompleteSystem struct {
	strToSearch *strings.Builder
	prefixDB map[string]*PrefixDB
}

type PrefixDB struct {
	sentences []string
	hotSentences []*HotSentence
	hotSentence *HotSentence
}

type HotSentence struct {
	word string
	hot int
}

func constructor(sentences []string, times []int) autocompleteSystem {
	a := autocompleteSystem{strToSearch: new(strings.Builder)}
	hotSentences := buildHotSentences(sentences, times)
	sortHotSentences(hotSentences)
	a.buildPrefixDB(hotSentences)
	return a
}

func (a *autocompleteSystem) Input(c byte) []string {
	if c == '#' {
		a.updatePrefixDB(a.strToSearch.String())
		a.strToSearch.Reset()
		return []string{}
	}

	a.strToSearch.WriteByte(c)
	if db, ok := a.prefixDB[a.strToSearch.String()]; ok {
		return db.sentences
	}

	return []string{}
}

func buildHotSentences(sentences []string, times []int) []*HotSentence {
	hotSentences := make([]*HotSentence, len(sentences))
	for i := range sentences {
		hotSentences[i] = new(HotSentence)
		hotSentences[i].word = sentences[i]
		hotSentences[i].hot = times[i]
	}
	return hotSentences
}

func sortHotSentences(hotSentences []*HotSentence) {
	sort.Slice(hotSentences, func(i, j int) bool {
		if hotSentences[i].hot == hotSentences[j].hot {
			return hotSentences[i].word < hotSentences[j].word
		}
		return hotSentences[i].hot > hotSentences[j].hot
	})
}

func (a *autocompleteSystem) buildPrefixDB(hotSentences []*HotSentence) {
	prefixDB := make(map[string]*PrefixDB)

	seen := make(map[string]map[*HotSentence]struct{})
	for i := range hotSentences {
		for j := 1; j < len(hotSentences[i].word)+1; j++  {
			prefix := hotSentences[i].word[:j]

			if seen[prefix] == nil {
				seen[prefix] = make(map[*HotSentence]struct{})
			}

			if _, ok := seen[prefix][hotSentences[i]]; !ok  {
				seen[prefix][hotSentences[i]] = struct{}{}
				if _, found := prefixDB[prefix]; !found  {
					prefixDB[prefix] = new(PrefixDB)
				}
				prefixDB[prefix].hotSentences = append(prefixDB[prefix].hotSentences, hotSentences[i])

				if len(prefixDB[prefix].sentences) < 3 {
					prefixDB[prefix].sentences = append(prefixDB[prefix].sentences, hotSentences[i].word)
				}
			}
		}
		prefixDB[hotSentences[i].word].hotSentence = hotSentences[i]
	}

	a.prefixDB = prefixDB
}

func (a *autocompleteSystem) updatePrefixDB(word string) {
	var newHotSentence *HotSentence
	if _, ok := a.prefixDB[word]; ok {
		if a.prefixDB[word].hotSentence == nil {
			newHotSentence = new(HotSentence)
			newHotSentence.word = word
			newHotSentence.hot = 1
		} else {
			a.prefixDB[word].hotSentence.hot++
		}
	} else {
		newHotSentence = new(HotSentence)
		newHotSentence.word = word
		newHotSentence.hot = 1
	}


	for j := 1; j < len(word)+1; j++  {
		prefix := word[:j]
		if _, ok := a.prefixDB[prefix]; !ok  {
			a.prefixDB[prefix] = new(PrefixDB)
			a.prefixDB[prefix].hotSentences = append(a.prefixDB[prefix].hotSentences, newHotSentence)
			a.prefixDB[prefix].sentences = append(a.prefixDB[prefix].sentences, word)
		} else {
			if newHotSentence != nil {
				a.prefixDB[prefix].hotSentences = append(a.prefixDB[prefix].hotSentences, newHotSentence)
			}
			sortHotSentences(a.prefixDB[prefix].hotSentences)

			sentences := make([]string, 0)
			for i := 0; i < 3 && i < len(a.prefixDB[prefix].hotSentences); i++ {
				sentences = append(sentences, a.prefixDB[prefix].hotSentences[i].word)
			}

			a.prefixDB[prefix].sentences = sentences
		}
	}

	if newHotSentence != nil {
		a.prefixDB[word].hotSentence = newHotSentence
	}
}
