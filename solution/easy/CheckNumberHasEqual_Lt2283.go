package easy

func digitCount(num string) bool {
	cnt := map[rune]int{}
	for _, c := range num {
		cnt[c-'0']++
	}
	for i, c := range num {
		if cnt[rune(i)] != int(c-'0') {
			return false
		}
	}
	return true
}
