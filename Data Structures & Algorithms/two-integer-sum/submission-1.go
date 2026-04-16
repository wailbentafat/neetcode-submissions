func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if value, ok := m[diff]; ok {
			return []int{value, i}
		}

		m[nums[i]] = i
	}

	return nil
}