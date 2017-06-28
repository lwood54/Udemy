// 6. Find a problem at projecteuler.net then create the solution. Add a comment beneath
//     your solution that includes a description of the problem. Upload your solution to github.
//     Tweet me a link to your solution.

/////// PROJECT EULER PROBLEM NUMBER 10 ///////
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(sumOfPrimes(10)) // should be 17 (2 + 3 + 5 + 7)
	fmt.Println(sumOfPrimes(2000000))
}

func sumOfPrimes(upperLimit int) int {
	var total int
	if upperLimit > 2 {
		total = 2
	} else {
		return 2
	}
	for i := 3; i < upperLimit; i += 2 {
		if isPrime(i) {
			total += i
		}
	}
	return total
}

func isPrime(num int) bool {
	numSqrt := int(math.Sqrt(float64(num)))
	if num%2 != 0 {
		for i := 3; i <= numSqrt; i += 2 {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
