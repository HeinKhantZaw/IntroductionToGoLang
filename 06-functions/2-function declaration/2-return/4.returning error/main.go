package main

import (
	"errors"
	"fmt"
)

func main() {
	//function return with error
	//num1, num2 := 10, 13
	num1, num2 := 10, 10
	result1, result2, err := swap(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1, result2)
	}

}

//swap() return two integers
func swap(num1, num2 int) (int, int, error) {
	if num1 == num2 {
		err := errors.New("Can't swap if the numbers are equal")
		return num2, num1, err
	} else {
		return num2, num1, nil
	}
}
