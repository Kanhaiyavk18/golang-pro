package main

import (
	"fmt"
	"strings"
)

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func CountWordsFunc(s string, f func(rune) bool) int {
	return len(strings.FieldsFunc(s, f))
}

func ExampleCountWords() {
	wc := CountWords("  Assignments on Go Programming Language \n")
	fmt.Println(wc)

}

func ExampleCountWordsFunc() {
	f := func(c rune) bool {
		return c == ','
	}
	wc := CountWordsFunc("Xoraint Solution Pvt ltd", f)
	fmt.Println(wc)

}
