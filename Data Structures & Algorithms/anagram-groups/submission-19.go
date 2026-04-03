func groupAnagrams(strs []string) [][]string {
    strsLen := len(strs)
    strsMap := make(map[string][]string, 0)

    for i := 0; i < strsLen; i++ {
        sortedStr := string(byteSort([]byte(strs[i])))
        strsMap[sortedStr] = append(strsMap[sortedStr], strs[i])
    }

    result := make([][]string, 0)
    for _, v := range strsMap {
        result = append(result, v)
    }
    return result
}

func byteSort(arr []byte) []byte {
    if len(arr) < 2 {
        return arr
    }
    arrLen := len(arr)
    piv := arr[arrLen / 2]
    right := []byte{}
    mid := []byte{}
    left := []byte{}

    for i := 0; i < arrLen; i++ {
        if piv < arr[i] {
            left = append(left, arr[i])    
        } else if piv == arr[i] {
            mid = append(mid, arr[i])
        } else {
            right = append(right, arr[i])   
        }
    }

    res := append(byteSort(left), mid...)    
    return append(res, byteSort(right)...)
}