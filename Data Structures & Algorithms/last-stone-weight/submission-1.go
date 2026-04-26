import "slices"
func lastStoneWeight(stones []int) int {
	slices.Sort(stones)
	length := len(stones)
	for length > 1{
		h_1 := stones[length - 1]
		h_2 := stones[length - 2]

		if h_1 == h_2 {
			length = length - 2
			continue
		}
		stones[length -2] = stones[length -1]
		length = length - 1
		stones[length - 1] = h_1 - h_2
		slices.Sort(stones)
	}
	if length == 1 {
		return stones[0]
	}
	return 0
}
