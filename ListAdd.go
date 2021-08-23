package main

import "fmt"
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
func print(head *ListNode) []int {
	res := make([]int,0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}


func ListAdd(h1 *ListNode, h2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	head1 := h1
	head2 := h2
	var node *ListNode

	for head1 != nil {
		stack1 = append(stack1,head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		stack2 = append(stack2,head2.Val)
		head2 = head2.Next
	}
	carry := 0
	for len(stack1) != 0 || len(stack2) != 0 || carry > 0 {
		sum := 0
		if len(stack1) != 0 {
			//POP stack1
			sum += stack1[len(stack1) - 1]
			stack1 = stack1[:len(stack1) - 1]
		}
		if len(stack2) != 0 {
			//POP stack2
			sum += stack2[len(stack2) - 1]
			stack2 = stack2[:len(stack2) - 1]
		}
		sum += carry
		tmp := &ListNode{Val: sum % 10}
		carry = sum / 10
		tmp.Next = node
		node = tmp

	}
	return node

}
func main() {
	h1 := &ListNode{Val: 2}
	h2 := &ListNode{Val: 4}
	h3 := &ListNode{Val: 3}
	h1.Next = h2
	h2.Next = h3
	h4 := &ListNode{Val: 5}
	h5 := &ListNode{Val: 6}
	h6 := &ListNode{Val: 4}
	h4.Next = h5
	h5.Next = h6
	hsum := ListAdd(h1,h4)
	res := print(hsum)
	fmt.Println(res)

}
