package medium

func countPrimes(n int) int {
	res := 0
	isPrimes := make([]bool, n)

	for i := 2; i < n; i++ {
		for j := i * i; j < n; j += i {
			if isPrimes[j] {
				break
			}
			isPrimes[j] = true
		}
	}

	for r := 2; r < n; r++ {
		if !isPrimes[r] {
			res++
		}
	}

	return res
}
