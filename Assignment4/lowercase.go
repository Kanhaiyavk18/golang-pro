// convert the given string to lowercase
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	str1 := "WelcomE, Xoriant Solution Pvt Ltd**"
	str2 := "Assignment oF Golang"
	str3 := "HELLO GOLANG"
	str4 := "lowercase conversion"

	fmt.Println("Strings (before):")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)
	fmt.Println("String 3:", str3)
	fmt.Println("String 4:", str4)

	res1 := strings.ToLower(str1)
	res2 := strings.ToLower(str2)
	res3 := strings.ToLower(str3)
	res4 := strings.ToLower(str4)

	fmt.Println("\nStrings (after):")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3:", res3)
	fmt.Println("Result 4:", res4)
}
