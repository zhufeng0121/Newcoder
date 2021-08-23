package main

import "fmt"

func sink(arr []int, index int, k int) {
	//这里的意义不仅仅在于只是为了交换而用，因为arr[index]会被重新赋值
	temp := arr[index]
	child := index << 1 + 1
	for child < k {
		if child + 1 < k && arr[child] < arr[child + 1] {
			child++
		}
		//这里的比较一定要用temp
		if temp > arr[child] {
			break
		}
		arr[index] = arr[child]
		index = child
		child = child << 1 + 1
	}
	arr[index] = temp
}


func HeapSort(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		sink(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		//最后一个元素和第一个元素进行交换
		arr[0], arr[i] = arr[i], arr[0]
		//下沉调整最大堆
		sink(arr, 0,i)
	}
}

func main() {
	arr := []int {4,5,1,6,2,7,3,8}
	HeapSort(arr)
	//GetLeastK(arr,3)
	fmt.Println(arr)
}


//获取最小的K个数，要建立容量为K的最大堆，对于arr 数组，首先可以
func GetLeastK(arr []int, k int){
	for i := len(arr) / 2 ; i >= 0; i-- {
		sink(arr,i,k)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			arr[0], arr[i] = arr[i], arr[0]
			sink(arr,0,k)
		}
	}
	fmt.Println(arr)

}


