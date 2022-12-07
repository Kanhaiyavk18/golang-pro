// find the factorial of any number Using Recursion
package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	fmt.Println(factorial(15))
}
