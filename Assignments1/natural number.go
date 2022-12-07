package main

import "fmt"

func sum(n int) int {

	if n == 0 {

		return 0
	} else {

		return n + sum(n-1)
	}
}

func main() {

	number := 25
	fmt.Println("Sum of first n numbers using recursion in Golang Program")
	fmt.Printf("Sum of the first %d numbers is %d", number, sum(number))

}
