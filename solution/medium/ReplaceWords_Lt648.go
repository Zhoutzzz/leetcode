package medium

import "strings"

func replaceWords(dictionary []string, sentence string) (res string) {
	split := strings.Split(sentence, " ")
	for _, v := range split {
		for _, s := range dictionary {
			if !strings.HasPrefix(v, s) {
				continue
			} else {
				v = s
			}
		}
		res += v + " "
	}
	return res[:len(res)-1]
}

/*
字典树
*/
func replaceWords_tree(dictionary []string, sentence string) (res string) {
	type trie map[rune]trie
	root := trie{}
	for _, s := range dictionary {
		cur := root
		for _, v := range s {
			if cur[v] == nil {
				cur[v] = trie{}
			}
			cur = cur[v]
		}
		cur['#'] = trie{}
	}
	split := strings.Split(sentence, " ")
	for i, v := range split {
		cur := root
		for j, c := range v {
			if cur['#'] != nil {
				split[i] = v[:j]
				break
			} else if cur[c] == nil {
				break
			} else {
				cur = cur[c]
			}
		}
	}
	return strings.Join(split, " ")
}
