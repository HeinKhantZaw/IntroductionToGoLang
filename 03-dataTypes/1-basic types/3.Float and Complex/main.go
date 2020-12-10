package main

import "fmt"

func main() {
	var myFloat32 float32 = 4.5
	var myFloat float64 = 9.12 //The default type for float will be 'float64'

	var myComplex64 complex64 = 1 + 4i
	var myComplex complex128 = 10 + 11i //The default type for complex will be 'complex128'

	fmt.Printf("Type: %T | Value: %v\n", myFloat32, myFloat32)
	fmt.Printf("Type: %T | Value: %v\n", myFloat, myFloat)
	fmt.Printf("Type: %T | Value: %v\n", myComplex64, myComplex64)
	fmt.Printf("Type: %T | Value: %v\n", myComplex, myComplex)
}
