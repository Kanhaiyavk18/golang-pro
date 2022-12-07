// Golang program to reverse a number
package main

import "fmt"

func reverseNumber(num int) int {

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func main() {
	fmt.Println(reverseNumber(9876))
	fmt.Println(reverseNumber(576))
	fmt.Println(reverseNumber(12345))
}
