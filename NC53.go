package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 删除链表的倒数第N个节点
// 双重指针： 第一个指针先走N步，然后再一起走，当到达末尾时，第二个指针指向了要删除节点的前一个节点
// 要考虑删除的是头结点的情况

//快慢指针  需要构造虚拟头结点处理边界
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}
	newHead := &ListNode{}
	newHead.Next = head
	fast, slow := newHead,newHead
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return newHead.Next
}





