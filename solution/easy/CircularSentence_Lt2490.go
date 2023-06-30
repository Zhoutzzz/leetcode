package easy

func isCircularSentence(sentence string) bool {
	n := len(sentence)
	if sentence[n-1] != sentence[0] {
		return false
	}
	for i := 0; i < n; i++ {
		if sentence[i] == ' ' && sentence[i+1] != sentence[i-1] {
			return false
		}
	}
	return true
}
