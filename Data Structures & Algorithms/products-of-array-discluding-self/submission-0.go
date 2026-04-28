func productExceptSelf(nums []int) []int {
	n := len(nums)
	tempArray := make([]int, n)
	tempArrayReverse := make([]int, n)
	output := make([]int, n)
	tempArray[0] = 1
	for i := 1; i < n; i++ {
		tempArray[i] = tempArray[i-1] * nums[i-1]
	}
	tempArrayReverse[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		tempArrayReverse[i] = tempArrayReverse[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		output[i] = tempArray[i] * tempArrayReverse[i]
	}

	return output
}