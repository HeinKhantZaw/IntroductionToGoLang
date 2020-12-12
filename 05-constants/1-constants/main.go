package main

import "fmt"

func main() {

	const Pi = 3.14

	const (
		squareRootOf2 = 1.41421
		squareRootOf3 = 1.73205
		squareRootOf5 = 2.23606
	)
	fmt.Println("π = ", Pi)
	fmt.Println("√2 * √3 * √5  = ", squareRootOf2*squareRootOf3*squareRootOf5)
}
