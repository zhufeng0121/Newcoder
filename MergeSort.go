package main

func mergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := start + (end - start) >> 1
	mergeSort(arr,start,mid)
	mergeSort(arr,mid + 1,end)
	mergeArray(arr,start,mid,end)

}

func mergeArray(arr []int,start, mid, end int) {
	tmp := make([]int, end- start + 1)
	p1 := start
	p2 := mid + 1
	p := 0
	for p1 <= mid && p2 <= end {
		if arr[p1] < arr[p2] {
			tmp[p] = arr[p1]
			p1++
		} else {
			tmp[p] = arr[p2]
			p2++
		}
		p++
	}
	for p1 <= mid {
		tmp[p] = arr[p1]
		p1++
		p++
	}
	for p2 <= end {
		tmp[p] = arr[p2]
		p2++
		p++
	}
	for i := 0; i < len(tmp); i++ {
		arr[i +start] = tmp[i]
	}

}