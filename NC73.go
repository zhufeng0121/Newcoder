package main

//数组中出现次数超过一半的次数

func MoreThanHalfNum_Solution( numbers []int ) int {
	res, count := numbers[0], 1
	for i := 1; i < len(numbers); i++ {
		if count > 0 {
			if res == numbers[i] {
				count++
			} else {
				count--
			}
		} else {
			res = numbers[i]
			count++
		}
	}
	return res
}


