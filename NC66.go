package main


func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	p1,p2 := pHead1, pHead2

	for p1 != p2{
		if p1 == nil {
			//这里回头的时候要注意 pHead2
			p1 = pHead2
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			//这里回头的时候要注意 pHead1
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}
	return p1
}