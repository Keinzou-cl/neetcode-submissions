func twoSum(nums []int, target int) []int {
    storage := make(map[int]int)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		val, ok := storage[difference]
		if ok && (nums[i] + difference == target) {
			result = append(result, val, i)
		}
		storage[nums[i]] = i 
	}
	return result
}
