package main

/**

给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

 */
func isValid( s string ) bool {
	stack := make([]byte,0)
	for i := 0 ;i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack,s[i])
		} else if len(stack) != 0 && s[i] == ')' && stack[len(stack) -1] == '(' ||
			s[i] == ']'&& len(stack) != 0  && stack[len(stack) -1] == '[' ||
			len(stack) != 0 && s[i] == '}' && stack[len(stack) -1] == '{' {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}