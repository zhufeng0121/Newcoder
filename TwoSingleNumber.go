package main

//数组中有两个数出现一次，其余的出现两次，返回这两个数

func singleNumber(nums []int) []int {
	res, div := 0, 1
	for _, v := range nums {
		res ^= v
	}
	for res & div != 1 {
		div <<= 1

	}
	a,b := 0, 0
	for _, v := range nums {
		if v & div == 0 {
			a ^= v
		}else {
			b ^= v
		}
	}
	return []int{a,b}
}