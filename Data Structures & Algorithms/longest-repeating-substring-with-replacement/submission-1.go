func characterReplacement(s string, k int) int {
count := [26]int{}

l:=0
maxf:=0
best:=0

for r:=0;r<len(s);r++{
	i:=s[r]-'A'
	count[i]++

	maxf = max(maxf, count[i])

	window:=r-l+1

	if window-maxf>k{
		count[s[l]-'A']--
		l++
		window -=1
	}
	best = max(best, window)
}
return best
}
