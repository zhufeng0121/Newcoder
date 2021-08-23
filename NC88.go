package main

import "fmt"

//根据快排思想找到第K小的数

//快排中有partition函数，思路应该就是利用partition函数进行切分
func findKth(a []int, n int, K int) int {
	return findK(a,0,n-1,K)

}
func findK(arr []int, left,right int, k int) int {
	i := partition(arr,left,right)
	if i + 1 == k {
		return arr[k]
	} else if i + 1 > k {
		return findK(arr,left, i -1, k)
	} else {
		return findK(arr,i+1,right,k)
	}

}

//返回已经排好序的元素的序号
func partition(a []int, start int, end int) int {
	//选取第一个元素作为pivot
	pivot := a[start]
	left, right := start, end
	for left != right {
		for left < right && a[right] <= pivot {
			right--
		}
		for left < right && a[left] > pivot {
			left++
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
		}

	}
	a[left], a[start] = a[start], a[left]
	return left
}

func main() {
	arr := []int{1332802,1177178,1514891,871248,753214,123866,1165786,93724,9868,7986}
	fmt.Println(findKth(arr,len(arr),4))
}