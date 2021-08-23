package main

//表达式求值

func solveI( s string ) int {
	operand, operator := make([]byte,0), make([]byte,0)

	for i := 0 ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			//如果元素是操作数，则入操作数的栈
			operator = append(operator,s[i])
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' {
			operand = append(operand,s[i])
		} else if s[i] == ')' {

		}
	}
}

