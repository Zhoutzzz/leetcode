package medium

import (
	"fmt"
	"strconv"
)

func FractionAddition(expression string) string {
	if expression[0] != '-' {
		expression = "+" + expression
	}

	numerator, denominator := 0, 1
	for i, v := range expression {
		if v == '/' {
			c1, c2 := "", ""
			for x := i + 1; x < len(expression); x++ {
				if x == len(expression)-1 || expression[x] == '+' || expression[x] == '-' {
					if x == len(expression)-1 {
						x += 1
					}
					c1 = expression[i+1 : x]
					break
				}
			}
			for x := i; x >= 0; x-- {
				if expression[x] == '+' || expression[x] == '-' {
					c2 = expression[x:i]
					break
				}
			}
			x, _ := strconv.Atoi(c1)
			y, _ := strconv.Atoi(c2)
			if numerator == 0 {
				denominator = x
				numerator = y
			} else {
				numerator = numerator*x + y*denominator
				denominator *= x
			}
		}
	}
	g := gcd(numerator, denominator)
	if g < 0 {
		g = g * -1
	}
	return fmt.Sprint(numerator/g) + "/" + fmt.Sprint(denominator/g)
}

func gcd(x, y int) int {
	v := x % y
	if v == 0 {
		return y
	}
	return gcd(y, v)
}
