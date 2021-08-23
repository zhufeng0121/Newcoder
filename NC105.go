package main

//寻找最左的二分查找

func search( nums []int ,  target int ) int {
	left, right := 0, len(nums) -1
	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			if mid == 0 || nums[mid -1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}
