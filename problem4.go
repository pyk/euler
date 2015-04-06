// solved: https://projecteuler.net/problem=4
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// first we find all palindromics number
	var palindromics []int64
	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				palindromic, err := strconv.ParseInt(fmt.Sprintf("%d%d%d%d%d%d", a, b, c, c, b, a), 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				palindromics = append(palindromics, palindromic)
			}
		}
	}

	// find the largest palindromic number from the product of 3-digit numbers
	var palindromic int64
	for _, p := range palindromics {
		for i := int64(100); i < 1000; i++ {
			if p%i == 0 {
				result := p / i
				if result <= 999 {
					palindromic = p
				}
			}
		}
	}

	fmt.Println(palindromic)
}
