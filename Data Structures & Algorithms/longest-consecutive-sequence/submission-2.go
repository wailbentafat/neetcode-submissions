func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	arrayMap := make(map[int]int)
	var startpoints []int

	for i, j := range nums {
		_, okee := arrayMap[j]
		if okee {
			continue
		}
		arrayMap[j] = i
	}

	for _, j := range nums {
		_, ok := arrayMap[j-1]
		if !ok {
			_, found := arrayMap[j+1]
			if found {
				startpoints = append(startpoints, j)
			}
		}
	}

	output := 1

	for _, start := range startpoints {
		current := 1
		value := start

		for {
			_, oke := arrayMap[value+1]
			if !oke {
				break
			}
			current++
			value++
		}

		if current > output {
			output = current
		}
	}

	return output
}