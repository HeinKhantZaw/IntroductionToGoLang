package main

import "fmt"

func main() {

	// for is the only looping construct in Go

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// writing while loop in Go
	sum := 1
	for sum < 1000 {
		fmt.Print("\n", sum, "+", sum, "=")
		sum += sum
		fmt.Println(sum)
	}

}
