func topKFrequent(nums []int, k int) []int {

freq := make(map[int]int)
for _, num := range nums{
	freq[num]++
}

buckets := make([][]int, len(nums)+1)

for num, count := range freq{
	buckets[count] = append(buckets[count], num)
}

res := make([]int, 0 , k)

for b:=len(buckets)-1;b>=0;b--{
	for _, num := range buckets[b]{
		res = append(res, num)
		if len(res)==k{return res}
	}
}
return res



}
