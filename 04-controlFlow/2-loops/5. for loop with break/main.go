package main

import "fmt"

func main() {
	for num := 1; num <= 100; num++ {
		if num%2 == 0 && num%3 == 0 {
			fmt.Printf("Least Common Divisor(LCD) of 2 and 3 is %d\n", num)
			break //exit the loop
		}
	}
}
