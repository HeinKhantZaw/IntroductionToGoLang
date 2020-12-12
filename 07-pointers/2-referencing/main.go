package main

import "fmt"

func main() {
	var a = 2020
	fmt.Println("Value:", *&a, "Pointer:", &a)
	// Ampersand is used to get a pointer to the memory address where var a is stored
	// Asterisk is used to get the value of the memory address that the pointer was pointing at
	// "*&a" is same as "a"
}
