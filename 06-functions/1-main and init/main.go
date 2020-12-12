package main

import "fmt"

func main() {
	// main is the entry point of the executable programs
	fmt.Println("This is main function")
}

/*
	init() function is executed before the main() function call,
	so it does not depend to main() function.
	The main purpose of the init() function is to initialize the global variables
	that cannot be initialized in the global context.
*/
func init() {
	fmt.Println("Hello! This is init function.")
}

// You are allowed to create multiple init() function in the same program
// They execute in the order they are created
func init() {
	fmt.Println("Hi again. This is second init function!")
}
