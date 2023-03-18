package medium

func checkPalindromeFormation(a, b string) bool {
	return checkConcatenation(a, b) || checkConcatenation(b, a)
}

func checkConcatenation(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	if left >= right {
		return true
	}
	return checkSelfPalindrome(a[left:right+1]) || checkSelfPalindrome(b[left:right+1])
}

func checkSelfPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left >= right
}
