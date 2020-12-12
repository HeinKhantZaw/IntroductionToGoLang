package main

import "fmt"

func main() {

	// iota is just like an enumerated type in other languages like C or Java
	// each constant is assigned to an integer

	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	const (
		d = iota // 0
		e = iota // 1
		f = iota // 2
	)
	fmt.Println(a, b, c, d, e, f)

	const (
		Monday int = iota // writing like this also works too!
		Tuesday
		Wednesday
		Thursday
		Friday
	)
	fmt.Println("Day", Monday, "=", "Monday")
	fmt.Println("Day", Tuesday, "=", "Tuesday")
	fmt.Println("Day", Wednesday, "=", "Wednesday")
	fmt.Println("Day", Thursday, "=", "Thursday")
	fmt.Println("Day", Friday, "=", "Friday")

	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10) same as 1 times 2, 10 times (2^10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
		GB = 1 << (iota * 10) // 1 << (3 * 10)
		TB = 1 << (iota * 10) // 1 << (4 * 10)
	)
	fmt.Printf("KB=%d\nMB=%d\n", KB, MB, GB, TB)
}
