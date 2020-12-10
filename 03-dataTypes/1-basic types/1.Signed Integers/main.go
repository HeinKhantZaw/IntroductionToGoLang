package main

import "fmt"

func main() {
	var (
		myInt8  int8  = 127                 // 8-bit signed integer (-128 to 127)
		myInt16 int16 = 32767               //16-bit signed integer (-32768 to 32767)
		myInt32 int32 = 214748364           //32-bit signed integer (-2147483648 to 214748364)
		myInt64 int64 = 9223372036854775807 //64-bit signed integer (-9223372036854775808 to 9223372036854775807)
	)
	/*
		When you need an integer value you should use int
		unless you have a specific reason to use a sized or unsigned integer type.
	*/

	var myInt = 12 //The default type for integers will be 'int'

	// Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myHexNumber = 0x14 //Equivalent to 20 in decimal

	// Use prefix '0' for declaring octal numbers
	var myOctalNumber = 044 //Equivalent to 36 in decimal

	fmt.Printf("Type: %T | Value: %v\n", myInt, myInt)
	fmt.Printf("Type: %T | Value: %v\n", myInt8, myInt8)
	fmt.Printf("Type: %T | Value: %v\n", myInt16, myInt16)
	fmt.Printf("Type: %T | Value: %v\n", myInt32, myInt32)
	fmt.Printf("Type: %T | Value: %v\n", myInt64, myInt64)
	fmt.Printf("Type: %T | Value: %v\n", myHexNumber, myHexNumber)
	fmt.Printf("Type: %T | Value: %v\n", myOctalNumber, myOctalNumber)
}
