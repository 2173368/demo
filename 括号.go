package main

import "fmt"

func maie() {
	fmt.Println(isValid("[[]]"))
}
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]string, 0)
	for _, value := range s {
		strChar := string(value)
		// 如果是左括号，入栈
		if isLeftBracket(strChar) {
			stack = append(stack, strChar)
			continue
		}
		// 如果是右括号
		if len(stack) == 0 {
			return false
		}
		// 判断栈顶元素与映射关系是否一致
		top := stack[len(stack)-1]
		expected := getRelationship(top)

		if strChar == expected {
			//一致出栈
			stack = stack[:len(stack)-1]
		} else {
			// 不匹配，直接返回false
			return false
		}
	}
	//字符串遍历完 查看切片是否有剩余元素
	return len(stack) == 0
}

func isLeftBracket(s string) bool {
	return s == "(" || s == "[" || s == "{"
}

func getRelationship(s string) string {
	switch s {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	default:
		return ""
	}
}
