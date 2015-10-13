package src_test

import "testing"
import "."
import "fmt"

func TestInput(t *testing.T) {
	if false {
		t.Error("False")
	}
}

func ExampleCodeOne() {
	ints := [4]int{1, 7, 16383, 16384}

	encoded := src.Encode(ints[:])
	fmt.Println(encoded)
	// Output: [1 7 255 127 128 128 1]
}

func ExampleDecodeOne() {
	var bytes [3]byte = [3]byte{0x01, 0x05, 0x12}
	decoded := src.Decode(bytes[:])
	fmt.Println(decoded)
	// Output: [1 5 18]
}

func ExampleEncodeDecodeSmallInts() {
	ints := [7]int{3, 4, 5, 6, 16383, 16384, 7}
	fmt.Println(src.Decode((src.Encode(ints[:]))))
	// Output: [3 4 5 6 255 127 128 128 1 7]
}
