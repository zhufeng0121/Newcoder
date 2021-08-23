package main

//利用双栈实现最小栈
//一个栈存储元素，一个栈存储当前最小元素
type MinStack struct {
	Stack []int
	Min   []int
}

func (m *MinStack) push(x int) {
	m.Stack = append(m.Stack,x)
	if len(m.Min) == 0 {
		m.Min = append(m.Min,x)
	} else {
		top := m.Min[len(m.Min) - 1]
		if x <= top {
			m.Min = append(m.Min,x)
		}
	}
}

func (m *MinStack) pop() {
	if len(m.Stack) != 0 {
		tmp := m.Stack[len(m.Stack) - 1]
		m.Stack = m.Stack[:len(m.Stack) - 1]
		if tmp == m.Min[len(m.Min) - 1] {
			m.Stack = m.Stack[:len(m.Stack) - 1]

		}
	}

}

func (m *MinStack) top() int{
	if len(m.Stack) == 0 {
		return -1
	}
	return m.Stack[len(m.Stack) - 1]
}

func (m *MinStack) getMin() int {
	if len(m.Stack) == 0 {
		return -1
	}
	return m.Min[len(m.Min) - 1]

}



