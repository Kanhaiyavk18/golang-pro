package main

import "fmt"

func vowelsconsonantsCount(sentence string) (int, int) {

	var vowelsCount, consonantCount int

	vowelsCount = 0

	consonantCount = 0

	for i := 0; i < len(sentence); i++ {

		if sentence[i] == ' ' {
			continue
		}

		if sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] == 'i' || sentence[i] == 'o' || sentence[i] == 'u' ||
			sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] == 'I' || sentence[i] == 'O' || sentence[i] == 'U' {

			vowelsCount++
		} else {

			consonantCount++
		}
	}

	return vowelsCount, consonantCount
}
func main() {

	var sentence string
	var vowelsCount, consonantCount int

	sentence = "Xoriant is a Silicon Valley based product engineering, software development and technology services firm"
	fmt.Println("Program to find the count of vowels and consonants in the separate function.")

	vowelsCount, consonantCount = vowelsconsonantsCount(sentence)
	fmt.Println("Sentence:-\n", sentence)

	// printing the count of vowels and consonants
	fmt.Println("Result:- \nThe total number of vowels in the above sentence are", vowelsCount)
	fmt.Println("The total number of consonants in the above sentence are", consonantCount)
}
