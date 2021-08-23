package main

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int ) int {
	if len(rotateArray) == 0 {
		return 0
	}
	start := 0
	end := len(rotateArray) - 1

	for start <= end {
		mid := start + (end - start) >> 1
		if rotateArray[mid] < rotateArray[mid -1] {
			return mid
		} else if rotateArray[mid] > start {
			start = mid + 1

		} else if rotateArray[mid] < start {
			end = mid - 1
		}
	}
	return -1

}

