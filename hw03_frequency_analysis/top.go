package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordsCount struct {
	word  string
	count int
}

func Top10(text string) []string {
	woCoMap := make(map[string]int)
	woCoSls := make([]wordsCount, 0, len(woCoMap))
	var ans, unSignedText []string
	if len(text) == 0 {
		return ans
	}
	re := regexp.MustCompile(`[,.!?:;"'()\s]+-*\s*`)
	unSignedText = re.Split(text, -1)
	for _, word := range unSignedText {
		woCoMap[strings.ToLower(word)]++
	}
	for word, count := range woCoMap {
		woCoSls = append(woCoSls, wordsCount{word, count})
	}
	sort.Slice(woCoSls, func(i, j int) bool {
		if woCoSls[i].count > woCoSls[j].count {
			return true
		}
		if woCoSls[i].count < woCoSls[j].count {
			return false
		}
		return woCoSls[i].word < woCoSls[j].word
	})
	for _, wc := range woCoSls {
		ans = append(ans, wc.word)
	}
	if len(ans) < 10 {
		return ans[:len(ans)]
	}
	return ans[:10]
}
