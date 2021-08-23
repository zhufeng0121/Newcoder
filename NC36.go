package main

//二分查找
func findLeft(target int, nums []int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}