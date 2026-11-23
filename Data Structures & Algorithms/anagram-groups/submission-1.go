func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	storage := make(map[[26]int][]string, 0)
	for _, s := range strs {
		count := [26]int{}
		for _, ch := range s {
			count[ch - 'a']++
		}
		storage[count] = append(storage[count], s)
	}
	for _, val := range storage {
		result = append(result, val)
	}
	return result
}
