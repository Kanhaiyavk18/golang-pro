package main

import "fmt"

func main() {
	var perfNum, perfMin, perfMax, perfsum int
	perfsum = 0
	fmt.Print("Enter the Minimum and Maximum limit of Perfect Numbers = ")
	fmt.Scanln(&perfMin, &perfMax)

	for perfNum = perfMin; perfNum <= perfMax; perfNum++ {
		perfsum = 0
		for i := 1; i < perfNum; i++ {
			if perfNum%i == 0 {
				perfsum = perfsum + i
			}
		}
		if perfNum == perfsum {
			fmt.Print(perfNum, "\t")
		}
	}
	fmt.Println()
}
