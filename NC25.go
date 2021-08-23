package main

//删除有序链表中的重复元素
func deleteDuplicates( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, pre := head.Next,head

	for cur != nil {
		if cur.Val != pre.Val {
			cur = cur.Next
			pre = pre.Next
		} else {
			pre.Next = cur.Next
			cur = cur.Next
		}
	}
	return head
}
