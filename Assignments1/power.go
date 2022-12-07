//Write a go program to find power of any number using recursion.

package main

import "fmt"

func POWER(num int, power int) int {
	var result int = 1
	if power != 0 {

		result = (num * POWER(num, power-1))
	}
	return result
}
func main() {
	fmt.Println("Golang Program to calculate the power using recursion")

	var base int = 50
	var power int = 2
	var result int

	result = POWER(base, power)

	fmt.Printf("%d to the power of %d is: %d\n", base, power, result)
}
