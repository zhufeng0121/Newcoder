package main

import "fmt"

/**
 这道题  漫画算法讲解的很好

                              中位数 == 目标值
                           是-              - 否
                          -                    -
                      查找成功               中位数在旋转值的左
                                           是-       -否
                         最左侧元素 <= 目标 < 中位数    中位数 < 目标 <= 最右侧元素
                         是-                    -否     是-                 -否
                         -                        -    -                     -
                    查找目标在中位数左侧      查找目标在中位数右侧                查找目标在中位数左侧
*/

//在旋转的二维数组中寻找目标元素
/*
1,2,3,4,5,6,7
7,1,2,3,4,5,6
6,7,1,2,3,4,5
5,6,7,1,2,3,4
4,5,6,7,1,2,3
3,4,5,6,7,1,2
2,3,4,5,6,7,1

所以可以得出结论，如果 arr[mid] > arr[start],旋转点一定在mid的右侧，根据这个条件来
判断旋转点和mid节点的相对位置

 */

func rotatedBinarySearch(arr []int, target int) int {
	start, end := 0, len(arr) - 1

	for start <= end {
		mid := start + (end - start) >> 1
		if arr[mid] == target {
			return mid
		} else if arr[mid] > arr[start] {
			//旋转点在mid的右侧
			if arr[mid] > target && target >= arr[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			//旋转点在mid的左侧
			if arr[mid] < target && target <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1

			}
		}

	}
	return -1

}

func main() {
	arr := []int{9,10,11,12,13,1,3,4,5,8}
	fmt.Println(rotatedBinarySearch(arr,12))

}
