package main

func rotateString(A string, B string) bool {

	if len(A) == 0 && len(B) == 0 {
		return true
	}

	if len(A) == 0 || len(B) == 0 {
		return false
	}

	a := []byte(A)
	for i := 0; i < len(a); i++ {
		a = append(a[1:], a[0])
		if string(a) == B {
			return true
		}

	}
	return false
}
