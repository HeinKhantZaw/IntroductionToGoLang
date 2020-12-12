package main

import "fmt"

func main() {
	myName := "John Doe"
	welcome(myName) // When welcome() is called, you need to pass an argument
}

// function welcome() is declared with a parameter
func welcome(name string) {
	fmt.Println("Welcome", name)
}
