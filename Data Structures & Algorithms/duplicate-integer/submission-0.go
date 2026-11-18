func hasDuplicate(nums []int) bool {
    storage := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		storage[nums[i]] += 1
		if storage[nums[i]] > 1 {
			return true
		}
	}
	return false
}
