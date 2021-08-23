package main

//删除有序链表中的重复元素,保留只出现一次的元素

func deleteDuplicatesII( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//构造虚拟头结点
	vir := &ListNode{Next: head,Val: 0}

	slow, fast := head, head
	pre := vir

	for fast != nil {
		//注意这里要是 fast 的下一个节点的值和 slow 不同，也就是 fast 是重复链表段的最后一个节点
		for fast != nil && fast.Next != nil && fast.Next.Val == slow.Val {
			fast = fast.Next
		}
		if fast == slow {
			//当前节点和下面的节点不重复
			pre = fast
			slow = fast.Next
			fast = slow
		} else {
			//当前节点和下面的节点发生重复，
			//而且 fast 是重复节点的最后一个节点，所以将 pre->next = fast->next 就可以将重复节点全部删掉，
			//然后 slow 和 fast 均更新为 fast->next；
			pre.Next =fast.Next
			slow = pre.Next
			fast = slow
		}

	}
	return vir.Next
}