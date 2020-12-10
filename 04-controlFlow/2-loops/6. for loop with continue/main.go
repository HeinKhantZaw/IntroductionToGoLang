package main

import "fmt"

func main() {
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue //if num%2 == 0, skip execution of Printf() and continue the loop
		}
		fmt.Printf("%d ", num)
	}
}
