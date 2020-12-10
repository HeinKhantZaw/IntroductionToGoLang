package main

import "fmt"

func main() {
	var (
		strExample  string
		unsignedNum uint      //different uint types (uint8 uint16 uint32 uint64 uintptr)
		byteExample byte      // alias for uint8
		num         int       //different int types (int8  int16  int32  int64)
		doubleNum   float32   //float data types (float32 float64)
		complexNum  complex64 //Complex numbers with float32 real and imaginary parts (complex64 complex128)
		runeType    rune      //same as int32 and also represent Unicode code points.
		cond        bool
	)
	fmt.Printf("string:%s\nunsignedNum:%d\nbyteExample:%d\nnum:%d\n")
}
