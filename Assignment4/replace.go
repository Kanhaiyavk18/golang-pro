// Go program to illustrate how to
// replace characters in the given string
package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating and initializing strings
	str1 := "Welcome to Xoriant Solution Pvt ltd"
	str2 := "This is the best company in Pune"

	fmt.Println("Original strings")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)

	// Replacing strings
	// Using ReplaceAll() function
	res1 := strings.ReplaceAll(str1, "Xoriant", "XORIANT")
	res2 := strings.ReplaceAll(str2, "the", "THE")

	// Displaying the result
	fmt.Println("\nStrings(After replacement)")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)

}
