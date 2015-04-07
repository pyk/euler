// solved: https://projecteuler.net/problem=6
package main

import (
	"fmt"
)

func findTheSumOfSquare(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		square := i * i
		sum += square
	}
	return sum
}

func findTheSquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
func main() {
	a := findTheSumOfSquare(100)
	b := findTheSquareOfSum(100)
	fmt.Println(b - a)
}
