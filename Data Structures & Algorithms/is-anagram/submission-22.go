func isAnagram(s string, t string) bool {
    sLen := len(s)
    tLen := len(t)

    if sLen != tLen {
        return false
    }

    sMap := make(map[byte]int)
    for i := 0; i < sLen; i++ {
        sMap[s[i]]++
    }

    for i := 0; i < tLen; i++ {
        if _, k := sMap[t[i]]; !k {
            return false
        }
        sMap[t[i]]--
    }

    for _, v := range sMap {
        if v != 0 {
            return false
        }
    }

    return true
}
