package main

import "fmt"

func main() {
	num1, num2 := 10, 13
	fmt.Printf("When swap() is not called: %d %d\nAfter swap() is called: ", num1, num2)
	fmt.Println(swap(num1, num2))

}

//swap() return two integers
func swap(num1, num2 int) (int, int) {
	return num2, num1
}
