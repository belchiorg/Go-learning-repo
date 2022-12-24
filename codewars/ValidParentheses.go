package main

func ValidParentheses(parens string) bool {
	openCount := 0
	for _, char := range parens {
		if char == '(' {
			openCount++
		} else {
			openCount--
			if openCount < 0 {
				return false
			}
		}
	}
	if openCount == 0 {
		return true
	}
	return false
}
