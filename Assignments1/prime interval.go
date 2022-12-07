package main

import (
	"fmt"

	"github.com/afshin/prime"
)

func primeNumber(lower_limit, upper_limit uint64) {
	fmt.Printf("prime numbers between %d and %d are:\n", lower_limit, upper_limit)

	for i := lower_limit; i <= upper_limit; i++ {
		// getting bool value from isprime() function and using it as condition of if loop.
		if prime.IsPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	// initializing the variables to get the intervals from the user
	var lower_limit, upper_limit uint64

	// let us print prime numbers between 5 and 60.

	lower_limit = 5
	upper_limit = 60

	// calling the primeNumber() function
	primeNumber(lower_limit, upper_limit)
}
