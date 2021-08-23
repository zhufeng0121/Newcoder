package main

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func ReverseList( pHead *ListNode ) *ListNode {
	cur := pHead //当前节点
	var pre *ListNode
	for cur != nil {
		//post保存当前节点的下一个节点，当cur为最后一个节点，即cur.Next == nil时，
		//post 已经为 nil
		post := cur.Next
		//pre 赋值cur.Next
		cur.Next = pre
		//将prev 赋值为cur
		pre = cur
		// cur 赋值为post
		cur = post
	}
	return pre

}

//递归反转链表
func ReverseListRecur(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	newHead := ReverseListRecur(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return newHead

}