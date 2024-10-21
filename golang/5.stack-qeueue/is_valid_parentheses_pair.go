package main

func isValidParenthesesPair(str string) bool {
	stack := []string{}
	parenthesesMap := make(map[string]string)
	parenthesesMap["("] = ")"
	parenthesesMap["{"] = "}"
	parenthesesMap["["] = "]"
	isValid := true

	for _, char := range str {
		if _, ok := parenthesesMap[string(char)]; ok {
			stack = append(stack, string(char))
		} else {
			if len(stack) > 0 {
				if parenthesesMap[stack[len(stack)-1]] == string(char) {
					stack = stack[:len(stack)-1]
				} else {
					isValid = false
					break
				}
			} else {
				isValid = false
				break
			}
		}
	}

	if isValid && len(stack) > 0 {
		isValid = false
	}

	return isValid
}

// input => {([  ])}

func Problem1(str string) bool {
	isValid := true
	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	stk := []string{}

	for _, v := range str {

		if _, ok := m[string(v)]; ok {
			// if it's opening p
			stk = append(stk, string(v))
		} else {
			// closing  one is here
			if len(stk) > 0 {
				p := stk[len(stk)-1]
				if m[p] == string(v) {
					// we pop the element from the stk
					stk = stk[:len(stk)-1]
				} else {
					isValid = false
				}

			} else {
				isValid = false
			}
		}

	}

	return isValid
}
