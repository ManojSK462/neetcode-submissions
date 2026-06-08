func carFleet(target int, position []int, speed []int) int {

	idx := make([]int, len(position))
	for i:=0;i<len(position);i++{
		idx[i]=i
	}

	sort.Slice(idx, func(i, j int) bool{
		return position[idx[i]]>position[idx[j]]
	})
	worseTime:=0.0
	fleets:=0
	for _ , i := range idx{
		time := float64(target-position[i])/float64(speed[i])

		
		if time>worseTime{
			fleets++
			worseTime=time
			}

	}
	return fleets
}


