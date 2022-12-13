package easy

func checkIfPangram(sentence string) bool {
	exist := [26]bool{}
	for _, c := range sentence {
		exist[c-'a'] = true
	}
	for _, b := range exist {
		if !b {
			return false
		}
	}
	return true
}
