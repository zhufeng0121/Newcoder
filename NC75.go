package main

//数组中只出现一次的两个数字

func FindNumsAppearOnce( array []int ) []int {
	res, div := 0, 1
	for _, v := range array {
		res ^= v
	}
	for res & div == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range array {
		//这个判断条件一定要注意，到处都是坑....
		if v & div != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	if a > b {
		a, b = b, a
	}
	return []int{a,b}

}


