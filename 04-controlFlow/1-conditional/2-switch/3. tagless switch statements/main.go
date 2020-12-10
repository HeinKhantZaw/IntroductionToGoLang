package main

import "fmt"

func main() {

	// Actually, this is the short form of (if-else if) statement
	var age int
	fmt.Print("Enter your age:")

	fmt.Scanln(&age) // takes user input

	switch {
	case age <= 10:
		fmt.Println("You're a kid")
	case age > 10 && age < 20:
		fmt.Println("You're a teenager")
	case age >= 20 && age < 25:
		fmt.Println("You're a young adult")
	default:
		fmt.Println("You're a adult")
	}
}
