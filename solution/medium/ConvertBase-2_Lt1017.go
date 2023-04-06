package medium

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	bits := [32]int{}
	for i := 0; i < 32 && n != 0; i++ {
		if n&1 > 0 {
			bits[i]++
			if i&1 > 0 {
				bits[i+1]++
			}
		}
		n >>= 1
	}
	carry := 0
	for i := 0; i < 32; i++ {
		val := carry + bits[i]
		bits[i] = val & 1
		carry = (val - bits[i]) / -2
	}
	pos := 31
	res := []byte{}
	for pos >= 0 && bits[pos] == 0 {
		pos--
	}
	for pos >= 0 {
		res = append(res, byte(bits[pos])+'0')
		pos--
	}
	return string(res)
}
