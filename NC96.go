package main

//判断一个链表是否是回文结构
func isPail( head *ListNode ) bool {
	res := make([]int,0)
	for head != nil {
		res = append(res,head.Val)
		head = head.Next
	}
	for i := 0; i < len(res) >> 1; i++ {
		if res[i] != res[len(res) - 1 - i] {
			return false
		}
	}
	return true
}
