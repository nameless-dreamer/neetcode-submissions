func isPalindrome(s string) bool {
	str := toClearString(toLowerString(s))

	low := 0
	high := len(str) - 1
	for low <= high {
		if str[low] != str[high] {
			return false
		}
		low++
		high--
	}

	return true
}

func toLowerString(str string) string {
	strLen := len(str)
	newStr := ""
	for i := 0; i < strLen; i++ {
		value := str[i]
		if value >= 65 && value <= 90 {
			value += 32
		}
		newStr = newStr + string(value)
	}
	return newStr
}

func toClearString(str string) string {
	strLen := len(str)
	newStr := ""
	for i := 0; i < strLen; i++ {
		value := str[i]
		if value >= 48 && value <= 57 || value >= 97 && value <= 122 {
			newStr = newStr + string(value)
		}
	}
	return newStr
}