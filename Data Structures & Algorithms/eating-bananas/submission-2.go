func minEatingSpeed(piles []int, h int) int {
l:=1
r:=0


for _, num := range piles{
	if num>r {
		r=num
	}
}
res:=r
for l<=r{
	mid := (l+r)/2
	time:=0
	for i:=0;i<len(piles);i++{
		time += int(math.Ceil(float64(piles[i])/float64(mid)))
	}

	if time<=h{
		res = mid
		r = mid-1
	}else{
		l = mid+1
	}
}

return res


}




