// write is anagram function
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs{
		count := [26]int{} 	
		
		for j:=0;j<len(str);j++{
			count[str[j]-'a']++
		}
		groups[count] = append(groups[count], str)
	}


	res := [][]string{}
	for _,group := range groups{
		res = append(res, group)
	}

	return res
	
	
}
