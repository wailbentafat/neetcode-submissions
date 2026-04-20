type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var results string
    for _, str := range strs {
        results += fmt.Sprintf("%d#%s", len(str), str)
    }
    return results
}

func (s *Solution) Decode(encoded string) []string {
  var results []string
  i:=0
    j:=0
  for i <len(encoded){
   if encoded[i]=='#'{
        numberOp,_ :=strconv.Atoi(encoded[j:i])
        j=numberOp+i+1
        word:=encoded[i+1:j]
        results=append(results,word)
        i = j
        
    }else {
         i++
    }
  }
  return results
}
