package main

import "fmt"

func main() {
	// If Statement
	if true {
		fmt.Println("Printed")
	}
	var num = 20
	if num%5 == 0 {
		fmt.Println(num, "is divisible by 5")
	}

	// If with logical operators
	var age = 20
	if age >= 19 && age <= 21 {
		fmt.Println("My Age is between 19 and 21")
	}
	var a, b = true, false
	if a || b {
		fmt.Println("Will print if either a OR b is true")
	}
	// If with a short statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

}
