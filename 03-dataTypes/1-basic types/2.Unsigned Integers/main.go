package main

import "fmt"

func main() {
	var (
		myByte   byte   = 111                  // alias for uint8
		myUint8  uint8  = 255                  // 8-bit unsigned integer (0 to 255)
		myUint16 uint16 = 65535                //16-bit unsigned integer (0 to 65535)
		myUint32 uint32 = 4294967295           //32-bit unsigned integer (0 to 4294967295)
		myUint64 uint64 = 18446744073709551615 //64-bit unsigned integer (0 to 18446744073709551615)
	)
	var myUint uint = 2020 //unsigned integer

	fmt.Printf("Type: %T | Value: %v\n", myUint, myUint)
	fmt.Printf("Type: %T | Value: %v\n", myByte, myByte)
	fmt.Printf("Type: %T | Value: %v\n", myUint8, myUint8)
	fmt.Printf("Type: %T | Value: %v\n", myUint16, myUint16)
	fmt.Printf("Type: %T | Value: %v\n", myUint32, myUint32)
	fmt.Printf("Type: %T | Value: %v\n", myUint64, myUint64)
}
