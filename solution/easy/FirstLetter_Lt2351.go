package easy

func repeatedCharacter(s string) byte {
	seen := map[rune]bool{}
	for _, c := range s {
		if seen[c] {
			return byte(c)
		}
		seen[c] = true
	}
	return 0 // impossible
}
