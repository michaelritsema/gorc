package src

import "fmt"

//https://cwiki.apache.org/confluence/display/Hive/LanguageManual+ORC
// Byte Run Length Encoding
// doesn't handle runs of > 127
func startOfRun(bytes []byte) bool {
	if len(bytes) < 3 {
		return false;
	} else {
		return bytes[0] == bytes[1] && bytes[0] == bytes[2]
	}
}


func litSeq(bytes []byte) []byte {
	header := byte(len(bytes)*-1)
	return append([]byte{header},bytes...)
}

func runSeq(bytes []byte) []byte {
	header := byte(len(bytes))
	return append([]byte{header}, bytes[0])
}

func ByteRunLengthEncode(bytes []byte) []byte {
	fmt.Println("start of func")
	encodedBytes := []byte{}

	startIndex := -1

	for i:=0;i<len(bytes); {
		if startIndex == -1 {
			startIndex=i
		}

		if startOfRun(bytes[i:]) {
			if(i-1 - startIndex > 0) {
				//dumpLitSeq(bytes[startIndex:i])
				encodedBytes = append(encodedBytes, litSeq(bytes[startIndex:i])...)
			}

			startRun := i
			for x:=i;;x++ {
				if(x+1 >= len(bytes) || bytes[x] != bytes[x+1]) {
					//dumpRunSeq(bytes[startRun:x+1])
					encodedBytes = append(encodedBytes, runSeq(bytes[startRun:x+1])...)
					startIndex = -1
					i=x
					break;
				}
			}

		} else {
			if i==len(bytes) -1 {
				//dumpLitSeq(bytes[startIndex:])
				encodedBytes = append(encodedBytes, litSeq(bytes[startIndex:])...)
			}
		}
		i++

	}
	fmt.Println(encodedBytes)
	fmt.Println("end of func")
	return encodedBytes

}

func ByteRunLengthDecode
