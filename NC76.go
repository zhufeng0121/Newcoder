package main

//利用两个栈来实现队列

var stack1 [] int
var stack2 [] int

//入队 都压入第一个栈
func Push(node int) {
	stack1 = append(stack1,node)

}

//出队 判断stack2是否为空，如果不为空，则直接返回stack2的栈顶
//如果stack2为空，判断stack1是否为空，如果stack1为空，则直接返回-1，
//否则 将stack1中的元素倒入stack2中，并返回栈顶
func Pop() int{
	if len(stack2) == 0 {
		if len(stack1) == 0 {
			return -1
		} else {
			for i := len(stack1) - 1; i >= 0; i-- {
				stack2 = append(stack2,stack1[i])
				stack1 = stack1[:i]
			}
			res := stack2[len(stack2) - 1]
			stack2 = stack2[:len(stack2) -1]
			return res
		}

	} else {
		res := stack2[len(stack2) - 1]
		stack2 = stack2[:len(stack2) -1]
		return res

	}

}
