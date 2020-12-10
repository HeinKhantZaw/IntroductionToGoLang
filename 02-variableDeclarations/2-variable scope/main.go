package main

import "fmt"

var global = "visible within a package" //will be available in any function inside the main package
var sameName = "aaa"

func main() {
	var local = "visible inside main function"
	var sameName = "bbb"
	fmt.Println("var local:", local)
	fun1()
	fmt.Println("var sameName(inside Main):", sameName) // the output will be "bbb"
	/*
		A variable declared within an inner scope
		having the same name as variable declared
		in the outer scope will shadow the variable
		in the outer scope.
	*/
}
func fun1() {
	fmt.Println("var global:", global)
	fmt.Println("var sameName(outside Main):", sameName)
	//fmt.Println("var local:",local)	//this will give a compiler error
}
