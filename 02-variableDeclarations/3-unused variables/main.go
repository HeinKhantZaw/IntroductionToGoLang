package main

import "fmt"

func main() {

	// In go lang, we can't have variables that is declared but not used.

	//a := "Unused var"
	//if u remove comment, it will give an error because "var a" is unused
	b := "Used var"
	fmt.Println(b)
}
