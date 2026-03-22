package sliding_window

// Minimum Window Substring
// Given two strings s and t of lengths m and n respectively, 
// return the minimum window substring of s such that every character in t (including duplicates) 
// is included in the window. If there is no such substring, return the empty string "".

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	countT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	countS := make(map[byte]int)
	have, need := 0, len(countT)
	res := [2]int{-1, -1}
	resLen := 1 << 31 - 1
	l := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		countS[c]++
		if countT[c] > 0 && countS[c] == countT[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				resLen = r - l + 1
				res[0] = l
				res[1] = r
			}
			countS[s[l]]--
			if countT[s[l]] > 0 && countS[s[l]] < countT[s[l]] {
				have--
			}
			l++
		}
	}

	if resLen == 1 << 31 - 1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
