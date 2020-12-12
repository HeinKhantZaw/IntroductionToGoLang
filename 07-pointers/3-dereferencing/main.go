package main

import "fmt"

func main() {
	myInt := 123
	myPointer := &myInt
	fmt.Println("Before de-referencing myInt:", myInt)
	*myPointer = 1234 // change the value stored at myPointer
	fmt.Println("After de-referencing myInt:", myInt)

	//=========================================================================

	a := 43
	fmt.Println("Value of a:", a)
	fmt.Println("Pointer of a:", &a)

	b := &a // b and a now shares same address
	fmt.Println("Pointer of b:", b)
	fmt.Println("Value stored at b:", *b)

	*b = 42                       // The value at this address of b, change it to 42
	fmt.Println("Value of a:", a) // var a becomes 42 because b = &a
}
