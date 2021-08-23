package main

import "fmt"

//采用双指针 加 hashmap，该题和leetcode 3题思路相同

//思路如下:
//遍历数组，如果该数从来没有出现过，那么加入map中，同时更新长度和map  map用来存放每个元素 最后出现的位置(索引)
//如果该数出现过，那么从map中取出该index，如果该index在left和right的范围内，则left = index + 1 同时 map更新
//如果不在left和right的范围内，只更新长度即可
func maxLength( arr []int ) int {
	m := make(map[int]int)
	left := 0
	res := 0
	for i := 0; i < len(arr); i++ {
		index, ok := m[arr[i]]
		if !ok {
			res = maxInt(res, i - left + 1)
		} else {
			if index < left {
				res = maxInt(res, i - left + 1)
			} else {
				left = index + 1
			}
		}
		//map用来存放每个元素 最后出现的位置(索引)，因此不管一个元素是否出现，都要进行更新位置
		m[arr[i]] = i
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1,2,3,1,2,3,2,2}
	fmt.Println(maxLength(arr))
}
