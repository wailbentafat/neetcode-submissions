func hasDuplicate(nums []int) bool {

    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        val ,ok :=m[nums[i]]
 
        if !ok{
            m[nums[i]]=1
        }else{
            m[nums[i]]=val+1
            return true 
            
        }
        if val >1{
            return true
        }
    } 


    return false 
}