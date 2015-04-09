// solved: https://projecteuler.net/problem=7
package main

import (
	"fmt"
)

var (
	primes []int // the collection of prime number
)

// source: http://en.wikipedia.org/wiki/Primality_test#Go_implementation
func isPrime(value int) bool {
	if value <= 3 {
		return value > 1
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 0
	for {
		if isPrime(i) {
			primes = append(primes, i)
			if len(primes) == 10001 {
				fmt.Println(i)
				return
			}
		}
		i++
	}
}
