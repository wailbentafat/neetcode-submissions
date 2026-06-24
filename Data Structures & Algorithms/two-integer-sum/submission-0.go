func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	arr:=make ([]int , 20)
	sum:=0
	for i:=0 ;i<len(nums);i++{
		sum+=nums[i]
		m[nums[i]]= i 
	}
	for i:=0;i<len(nums);i++{
		diff :=target - nums[i]
		value ,ok := m[diff]
		if ok {
		return []int{i ,value}	
		}
	}
	return arr

}
