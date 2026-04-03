func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    s_table := make(map[byte]int, len(s))
    t_table := make(map[byte]int, len(t))
    for i := 0; i < len(s); i++ {
        s_table[s[i]]++
    }

    for i := 0; i < len(t); i++ {
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
