package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		syntax
		var variableName [ArrayLength] ElementType
	*/
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr)

	var arr1 [10]*float64
	floatNum := 4.5252
	arr1[0] = &floatNum
	fmt.Println(*arr1[0])

	var twoDimen [3][5]int // declaring two-dimensional array
	/*
	  Iterating two dimensional array in go
	*/
	for i,row :=range twoDimen{
		for j := range row{
			twoDimen[i][j] = rand.Intn(100) // assign random value between 0 and 100 to each index
		}
	}
	fmt.Println(twoDimen)
}
