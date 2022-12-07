package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hellooo GO... just GO ahead!"
	substr := "GO"
	count := strings.Count(str, substr)
	fmt.Println("The number of occurrences of substring in the string is: ", count)
}
