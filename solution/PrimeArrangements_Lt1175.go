package main

const mod int = 1e9 + 7

func numPrimeArrangements(n int) int {
	primeCount := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}
	return factorial(primeCount) * factorial(n-primeCount) % mod
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i % mod
	}

	return f
}
