func groupAnagrams(strs []string) [][]string {

    type matchingArr []string

    var result []matchingArr
    anagramMap := make(map[string]matchingArr) 

    for _, s := range strs {
        chars := []rune(s)
        for i := 1; i < len(chars); i++ {
            key := chars[i]
            j := i - 1
            for j >= 0 && chars[j] > key {
                chars[j+1] = chars[j]
                j--
            }
            chars[j+1] = key
        }
        sortedStr := string(chars)

        anagramMap[sortedStr] = append(anagramMap[sortedStr], s)
    }

    for _, group := range anagramMap {
        result = append(result, group)
    }

    finalResult := make([][]string, len(result))
    for i, v := range result {
        finalResult[i] = []string(v)
    }

    return finalResult
}