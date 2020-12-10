package main

import "fmt"

func main() {
	// no need to add break in switch
	var num = 3
	switch num {
	case 1:
		fmt.Println("No.1")
	case 2:
		fmt.Println("No.2")
	case 3:
		fmt.Println("No.3")
	case 4:
		fmt.Println("No.4")
	case 5:
		fmt.Println("No.5")
	default:
		fmt.Println("Default number")
	}
}
