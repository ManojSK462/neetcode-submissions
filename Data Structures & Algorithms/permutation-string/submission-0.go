func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	target := make([]int, 26)
	window := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		target[s1[i]-'a']++
	}

	l := 0

	for r := 0; r < len(s2); r++ {
		window[s2[r]-'a']++

		if r-l+1 > len(s1) {
			window[s2[l]-'a']--
			l++
		}

		if r-l+1 == len(s1) && sameCounts(target, window) {
			return true
		}
	}

	return false
}

func sameCounts(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}