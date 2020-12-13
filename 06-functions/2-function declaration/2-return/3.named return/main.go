package main

import "fmt"

func main() {
	num1, num2 := 10, 13
	fmt.Printf("When swap() is not called: %d %d\nAfter swap() is called: ", num1, num2)
	fmt.Println(swap(num1, num2))

}

// return values of a function can be named in GoLang
func swap(num1, num2 int) (swap1 int, swap2 int) {
	swap1 = num2
	swap2 = num1
	return swap1, swap2
}
