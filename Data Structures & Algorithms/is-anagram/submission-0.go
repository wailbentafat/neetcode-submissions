func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false 
    }

    m1:=make(map[byte]int)
    m2:=make(map[byte]int)
    for i:=0 ; i<len(s);i++ {
        val,ok:=m1[s[i]]
        if !ok{
            m1[s[i]] = 1
        }else {
            m1[s[i]]= val+1
        }
        val2,ok:=m2[t[i]]
        if !ok{
            m2[t[i]] = 1
        }else {
            m2[t[i]]= val2+1
        }
    }
    for i:=0 ; i<len(s);i++ {
        if m1[s[i]]!=m2[s[i]]{
            return false
        }
        if m1[t[i]]!=m2[t[i]]{
            return false
        }
    }

return true 


}
