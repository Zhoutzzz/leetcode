package medium

func generateParenthesis(n int) (res []string) {
	init := ""
	generateBracket(init, 0, 0, n, &res)
	return res
}

func generateBracket(curSeq string, open, close, max int, res *[]string) {
	if open+close == max<<1 {
		*res = append(*res, curSeq)
	}
	if open < max {
		curSeq += "("
		generateBracket(curSeq, open+1, close, max, res)
		curSeq = curSeq[:len(curSeq)-1]
	}
	if close < open {
		curSeq += ")"
		generateBracket(curSeq, open, close+1, max, res)
		curSeq = curSeq[:len(curSeq)-1]
	}
}
