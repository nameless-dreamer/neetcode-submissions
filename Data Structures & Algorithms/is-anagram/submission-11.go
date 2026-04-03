func isAnagram(s string, t string) bool {
    s_len := len(s)
    t_len := len(t)

    if s_len != t_len {
        return false
    }

    s_table := map[byte]int{}
    t_table := map[byte]int{}
    for i := 0; i < s_len; i++ {
        s_table[s[i]]++
    }

    for i := 0; i < t_len; i++ {
        if _, k := s_table[t[i]]; !k {
            return false
        }

        t_table[t[i]]++
    }

    for k, v := range t_table {
        if s_table[k] != v {
            return false
        }
    }

    return true
}
