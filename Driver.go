package main

import "./src"

func main() {
	bytes := []byte{2,2,4,4,4,1,1,1,1,2,3,4,5,8,8,8,8,8}
	src.ByteRunLengthEncode(bytes)
}