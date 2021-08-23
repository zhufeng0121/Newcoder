package main

import "fmt"

//掌握快速排序的思想
func quickSort(arr []int, start, end int ) {
	if start > end {
		return
	}
	pivot := patition(arr, start, end)
	quickSort(arr,start,pivot - 1)
	quickSort(arr,pivot + 1, end)
}

func patition(arr []int, start, end int) int {
	//选取第一个元素作为基准元素
	//pivot := arr[start]
	left ,right := start, end
	for left < right {
		for left < right && arr[right] > arr[start] {
			right--
		}
		for left < right && arr[left] <= arr[start] {
			left++
		}
		//交换left和right指向的元素
		if left < right {
			arr[left], arr[right] = arr[right],arr[left]
		}
	}
	//将pivot和指针重合点进行交换
	//按照以上写法，因为先走的是右指针，所以当左指针和右指针重叠时，重叠处一定是小于pivot，所以进行交换
	arr[start], arr[left] = arr[left], arr[start]
	return left
}

func main() {
	arr := []int{2,3,1,5,3,9,0}
	quickSort(arr, 0,len(arr) - 1)
	fmt.Println(arr)
}


