package main

import "fmt"

func main() {
	var age = 21
	if age < 10 {
		fmt.Println("You're a kid")
	} else if age >= 10 && age < 20 {
		fmt.Println("You're a teenager")
	} else if age >= 20 && age < 25 {
		fmt.Println("You're a young adult")
	} else {
		fmt.Println("You're a adult")
	}
}
