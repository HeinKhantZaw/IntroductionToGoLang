package main

import "fmt"

func main()  {

	//================================================================

	// Declaring a variable
	var num int
	num = 1
	fmt.Println("1. Number output:",num)

	var num1, num2 float32
	num1 = 20.1
	num2 = 3.1
	fmt.Println("2. Sum of two floats:",num1+num2)

	//================================================================

	// Declaring a variable with an initial value
	var number = 12
	fmt.Println("3. Number with initial value:",number)

	var stringInit = "4. String with initial value"
	fmt.Println(stringInit)

	//================================================================

	// Short Declaration
	str := "5. This is a string"
	fmt.Println(str)

	str1, year := "6. This is Year", 2020
	fmt.Printf("%s %d\n",str1,year)

	//================================================================

	// Multiple Declarations
	var firstName, lastName = "John","Doe"
	fmt.Println("7.",firstName,lastName)

	var (
		float float32
		cond bool
	)
	float = 3.2
	cond = true
	fmt.Println("8.",cond,float)

	var (
		name = "John"
		email = "john@gmail.com"
	)
	fmt.Printf("9. Name=%s email=%s\n",name,email)
}