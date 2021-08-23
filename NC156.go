package main

//数组中只出现一次的数，其他的数出现了K次

//思路是利用(bit)位操作

func foundOnceNumber( arr []int ,  k int ) int {
	count := make([]int, 32)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < 32; j++ {
			count[j] += arr[i] & 1
			arr[i] >>= 1
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= count[31 -i] % k
	}
	return res

}
