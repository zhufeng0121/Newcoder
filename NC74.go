package main

//统计一个数组在升序数组中出现的次数(此题)
func GetNumberOfK( data []int ,  k int ) int {
	//这里得注意一下判断，或者将两个函数返回不同的返回值也可以
	if FindFirst(data,k) == -1 {
		return 0
	}
	return FindLast(data,k) - FindFirst(data,k) + 1
}

//返回最左的
func FindFirst(data []int, target int ) int {
	start, end := 0, len(data) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if data[mid] > target {
			end = mid - 1
		} else if data[mid] < target {
			start = mid + 1
		} else {
			if mid == 0 || data[mid-1] != target {
				return mid
			}
			end = mid - 1
		}
	}
	return -1

}

func FindLast(data []int, target int) int {
	start, end := 0, len(data) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if data[mid] > target {
			end = mid - 1
		} else if data[mid] < target {
			start = mid + 1
		} else {
			if mid == len(data) - 1 || data[mid+1] != target {
				return mid
			}
			start = mid + 1
		}
	}
	return -1

}


