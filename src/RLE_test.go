package src_test

import "."
import "fmt"



func ExampleByteRunLengthEncode() {
	bytes := []byte{1, 1,1,1,1,2, 3 ,4 ,5, 6}
	encoded := src.ByteRunLengthEncode(bytes)
	fmt.Println(encoded)

}

