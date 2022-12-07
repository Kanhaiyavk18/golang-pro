package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	var ptr *[5]int

	ptr = &arr

	fmt.Println("Array elements: ")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", (*ptr)[i])
	}
}
