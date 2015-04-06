// solved: https://projecteuler.net/problem=2
package main

import (
	"fmt"
)

// Each new term in the Fibonacci sequence will be stored here
var (
	terms = []int{1, 2}
)

// findTerms find new fibonacci terms based on given number "a" and "b"
// this function will stop find new term if exceed given value "until"
func findTerms(a, b, until int) int {
	term := a + b
	if term > until {
		return 0
	}
	terms = append(terms, term)
	if term > 3 {
		b = a
	}
	return findTerms(term, b, until)
}

func main() {
	findTerms(1, 2, 4000000)
	sum := 0
	for _, v := range terms {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Println("Answer: ", sum)
}
