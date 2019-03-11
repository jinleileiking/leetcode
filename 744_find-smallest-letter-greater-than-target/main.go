package main

func nextGreatestLetter(letters []byte, target byte) byte {

	for _, v := range letters {
		if target < v {
			return v
		}
	}

	return letters[0]

}
