package easy

func f(n, m int) int {
	return (m + n/m*m) * (n / m) / 2
}

func sumOfMultiples(n int) int {
	return f(n, 3) + f(n, 5) + f(n, 7) - f(n, 3*5) - f(n, 3*7) - f(n, 5*7) + f(n, 3*5*7)
}
