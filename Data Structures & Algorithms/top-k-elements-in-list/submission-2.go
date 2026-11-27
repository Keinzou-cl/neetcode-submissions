func topKFrequent(nums []int, k int) []int {
	sort := make([][]int, len(nums) + 1)
	store := make(map[int]int)
	result := make([]int, 0)
	for i:= 0; i < len(nums); i++ { // {1:1, 2:2, 3:3}
		store[nums[i]]++ 
	}
	for key, count := range store {
		sort[count] = append(sort[count], key)
	}
	for j := len(sort) - 1; j >= 0; j-- {
		for _, val := range sort[j] {
			if len(result) == k {
				continue
			}
			result = append(result, val)
		}
	}
	return result
}
