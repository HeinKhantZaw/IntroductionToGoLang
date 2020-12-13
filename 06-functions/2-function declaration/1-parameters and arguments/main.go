package main

import "fmt"

func main() {
	myName := "John Doe"
	welcome(myName) // When welcome() is called, you need to pass an argument
	add(10, 2)
	showAge(myName, 17)
}

// function welcome() is declared with a parameter
func welcome(name string) {
	fmt.Println("Welcome", name)
}

//function sum() is declared with two parameters
func add(num1, num2 int) {
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}

func showAge(name string, age int) {
	fmt.Printf("%s is %d years old\n", name, age)
}
