package main

func isValid(s string) bool {
	// 切分字符串
	runes := make([]string, 0, len(s))
	stack := make([]string, len(runes))
	for _, value := range runes {
		if len(stack) > 0 {
			stack = append(stack, value)
			continue
		}
		// 判断栈顶元素与映射关系是否一致
		relationship := getRelationship(value)
		if stack[0] == relationship {
			//一致出栈
			stack = append(stack[:1], stack[2:]...)
			continue
		}
		//不一致入栈
		stack = append(stack[:1], append([]string{value}, stack[1:]...)...)
	}
	//字符串遍历完 查看切片是否有剩余元素
	if len(stack) > 0 {
		return false
	}
	return true
}

func getRelationship(s string) string {
	switch s {
	case "{":
		return "}"
	case "[":
		return "]"
	case "(":
		return ")"
	default:
		return ""
	}
}
