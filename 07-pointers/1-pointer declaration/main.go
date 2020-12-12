package main

import "fmt"

func main() {

	var intPointer *int // this variable holds a intPointer to an int
	var a = 20
	intPointer = &a     // & get the pointer of "var a" and assigned to intPointer
	var b = *intPointer // * get the value of intPointer and is assigned to "var b"
	fmt.Println("Value of a:", a, "Value of b:", b)
	fmt.Println("intPointer")
	fmt.Printf("Type:%T Address:%v Value:%d\n", intPointer, intPointer, *intPointer)

	//=============================================================================

	var stringPointer *string //this variable holds a pointer to a string
	myString := "Hello"
	stringPointer = &myString
	fmt.Println("stringPointer")
	fmt.Printf("Type:%T Address:%v Value:%s", stringPointer, stringPointer, *stringPointer)
}
