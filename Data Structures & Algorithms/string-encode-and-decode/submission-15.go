type Solution struct{}

func (s *Solution) Encode(strs []string) string {
 	str := ""
	for i := 0; i < len(strs); i++ {
		str = str + strs[i] + "]["
	}
	return str
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	str := ""
	for i := 0; i < len(encoded); i++ {
		if string(encoded[i]) == "]" {
			res = append(res, str)
			str = ""
			i++
			continue
		}
		str = str + string(encoded[i])
	}

	return res
}
