package main

import "fmt"

func main() {
	var myBool bool = true
	var myString string = "HELLO"
	var myByte byte = 'A'
	var myRune rune = '😂'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
	fmt.Printf("Type: %T | Value: %v\n", myBool, myBool)
	fmt.Printf("Type: %T | Value: %v\n", myString, myString)
	fmt.Printf("Type: %T | Value: %v\n", myByte, myByte)
	fmt.Printf("Type: %T | Value: %v\n", myRune, myRune)
}
