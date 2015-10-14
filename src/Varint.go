package src

// based on protocol buffers 128 variant
// compression for small numbers
// only supported unsigned
// https://cwiki.apache.org/confluence/display/Hive/LanguageManual+ORC

func VarintEncode(integers []int) []byte {
	bytes := []byte{}

	for _, v := range integers {

		for {
			nextOR := 0
			if v>>7 > 0 {
				nextOR = 128
			}

			bytes = append(bytes, byte(v&127|nextOR))
			if v>>7 == 0 {
				break
			}
			v = v >> 7
		}
	}
	return bytes
}

func VarintDecode(bytes []byte) []int {
	var integers []int
	seq := []byte{}

	for _, v := range bytes {

		seq = append(seq, v)

		if v&128 != 1 {
			// routine
			newInt := 0
			for i := len(seq) - 1; i >= 0; i-- {
				//fmt.Println("x")
				hiBit := 0
				if i > 0 {
					hiBit = int(seq[i-i] & 128)
				}
				newInt = newInt | (int((seq[i] | byte(hiBit))) & (0xFF << uint(i)))
			}

			integers = append(integers, newInt)
			seq = []byte{}
		}

	}

	return integers
}
