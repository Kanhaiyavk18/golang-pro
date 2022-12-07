package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Kanhaiya kumar\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
