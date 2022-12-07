// Convert the given string into uppercase
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	str1 := "WelcomE, Xoraint Solution pvt ltd"
	str2 := "Assignment of golang"
	str3 := "Hello Go world"
	str4 := "uppercase conversion"

	// Displaying strings
	fmt.Println("Strings (before):")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)
	fmt.Println("String 3:", str3)
	fmt.Println("String 4:", str4)

	// Converting all the string into uppercase
	// Using ToUpper() function
	res1 := strings.ToUpper(str1)
	res2 := strings.ToUpper(str2)
	res3 := strings.ToUpper(str3)
	res4 := strings.ToUpper(str4)

	// Displaying the results
	fmt.Println("\nStrings (after):")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3:", res3)
	fmt.Println("Result 4:", res4)
}
