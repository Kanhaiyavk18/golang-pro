// Comapare two string
package main

import (
	"fmt"
	"strings"
)

func main() {

	A := "Xoriant Solution Pvt Ltd"
	B := "Xoriant Solution Pvt Ltd"

	// using the Compare function
	if strings.Compare(A, B) == 0 {
		fmt.Println("Both the strings match.")
	} else {
		fmt.Println("The strings do not match.")
	}
}
