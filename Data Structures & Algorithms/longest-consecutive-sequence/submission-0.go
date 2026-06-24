func longestConsecutive(nums []int)int {
	arrayMap :=make(map[int]int)
	var startpoints int
	max :=nums[0]

		for i ,j :=range nums {
			if max <j{
				max=j
				println("max",max)
			}
			_,okee:=arrayMap[j]
			if okee{
				continue
			}else{
				print(j)
				arrayMap[j]=i
			}
		}
		for _,j :=range nums{
			_,ok:= arrayMap[j-1]
			if !ok{
				_,found:=arrayMap[j+1]
				if found{
					startpoints=j
				}
			}
		} 
		output :=1
		for start:=startpoints;start<=max;start++{
			_,oke:=arrayMap[start+1]
			if oke{
				output++
			}else{
				break
			}
		}
		return output
		
	}
