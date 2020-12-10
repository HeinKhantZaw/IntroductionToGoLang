package main

import "fmt"

func main() {
	var myBool bool = true
	var myNormalString string = "\t\"helloWorld\"" //Escape characters are interpreted
	var myRawString string = `\thello 
								World` //can contain line breaks(Escape characters are not interpreted)
	var myByte byte = 'A'
	var myRune rune = '😂'
	var ascii = 0x42 // numerical value = 66, character code = 'B'

	fmt.Printf("Type: %T | Value: %v\n", myBool, myBool)
	fmt.Printf("Type: %T | Value: %v\n", myNormalString, myNormalString)
	fmt.Printf("Type: %T | Value: %v\n", myRawString, myRawString)
	fmt.Printf("character : %c | numerical value: %v\n", myByte, myByte)
	fmt.Printf("character : %c | Unicode value: %U\n", myRune, myRune)
	fmt.Printf("character : %c | Hex : %#x\n", ascii, ascii)
}
