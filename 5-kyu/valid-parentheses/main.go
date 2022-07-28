package main

func ValidParentheses(parens string) bool {
	if len(parens)%2 == 1 {
		return false
	}

	res := 0
	for i := 0; i < len(parens); i++ {
		if parens[i] == '(' {
			res += 1
		} else {
			res -= 1
		}
		if res < 0 {
			return false
		}
	}
	if res == 0 {
		return true
	} else {
		return false
	}
}
