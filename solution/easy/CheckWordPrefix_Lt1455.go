package easy

import "strings"

func isPrefixOfWord(sentence string, searchWord string) (ans int) {
	split := strings.Split(sentence, " ")
	for i, w := range split {
		if strings.HasPrefix(w, searchWord) {
			return i + 1
		}
	}
	return -1
}

func isPrefixOfWord1(sentence string, searchWord string) (ans int) {
l1:
	for i, idx, n := 0, 1, len(sentence); i < n; i++ {
		start := i
		for i < n && sentence[i] != ' ' {
			i++
		}
		for i1 := range searchWord {
			if searchWord[i1] != sentence[start+i1] {
				idx++
				continue l1
			}
		}
		return idx
	}
	return -1
}
