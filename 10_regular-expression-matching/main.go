package main

func isMatch(s string, p string) bool {
	for k, _ := range p {
		// spew.Dump(p[k:], k)
		// if k == len(p)-1 {
		// 	return false
		// }
		if isMatchInner(s, p[k:]) {
			return true
		}
	}

	return false
}

func isMatchInner(s string, p string) bool {
	var lastChar byte
	var patternPos int
	for strPos := 0; strPos < len(s); {
		c := s[strPos]
		// spew.Dump(strPos)
		if patternPos > len(p)-1 {
			return false
		}

		// spew.Dump(string(c), string(p[patternPos]))
		// spew.Dump(strPos, "zzzzz")
		// spew.Dump(patternPos)
		if p[patternPos] != '.' && p[patternPos] != '*' {
			if c != p[patternPos] {
				return false
			}
			if lastChar != 0 {
				lastChar = 0
			}
			patternPos++
			strPos++
			continue
		}

		if p[patternPos] == '.' {
			strPos++
			patternPos++
			continue
		}

		if p[patternPos] == '*' {
			if lastChar == 0 {
				if patternPos == 0 {
					// "*xxxxxx"
					return false
				}
				lastChar = p[strPos-1]
			}
			if lastChar == '.' {
				// .* matches anything
				return true
			}
			if lastChar != c {
				// spew.Dump(string(lastChar), string(c))
				patternPos++
				lastChar = 0
				// return false
				continue
			}

			strPos++
			// spew.Dump(string(lastChar))
		}
	}

	return true
}
