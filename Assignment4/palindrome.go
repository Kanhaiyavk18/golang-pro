package main

import "fmt"

func checkPalindrome(num int) string {
	input_num := num
	var remainder int
	res := 0
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	if input_num == res {
		return "Palindrome"
	} else {
		return "Not a Palindrome"
	}
}

func main() {
	fmt.Println(checkPalindrome(151))
	fmt.Println(checkPalindrome(152))
	fmt.Println(checkPalindrome(1551))
	fmt.Println(checkPalindrome(1541))
}
