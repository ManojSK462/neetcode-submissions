func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := 0
	required := len(need)

	left := 0
	bestLen := len(s) + 1
	bestStart := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++

		if need[ch] > 0 && window[ch] == need[ch] {
			have++
		}

		for have == required {
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}

			leftChar := s[left]
			window[leftChar]--

			if need[leftChar] > 0 && window[leftChar] < need[leftChar] {
				have--
			}

			left++
		}
	}

	if bestLen == len(s)+1 {
		return ""
	}

	return s[bestStart : bestStart+bestLen]
}