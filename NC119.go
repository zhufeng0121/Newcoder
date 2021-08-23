package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */

//寻找K个最小的数，所以要建造容量为K的大顶堆
func GetLeastNumbers_Solution( input []int , k int) []int {
	//先去元数据中取K个元素放在长度为K的数组中，再把数组转换成最小堆
	//再依次取原数据中的K个之后的数据和堆的根节点，
	HeapMake(input, k)
	for i := k; i < len(input); i++ {
		if input[i] > input[0] {
			input[0] = input[i]
			down(input,0,k)
		}
	}
	return input[:k]

}


//利用nums进行建堆操作
func HeapMake(nums []int, k int) {
	//完全二叉树中只有数组下标小于或者等于(data.length) / 2 - 1的元素有孩子节点
	for i := (k - 2) /2 - 1; i >= 0; i-- {
		down(nums,i,k)
	}
}

func down(nums []int, index int, k int) {
	temp := nums[index]
	child := index << 1 + 1
	for child < k {
		//如果有右孩子，且右孩子小于左孩子的值,定位到右孩子节点
		if child + 1 < len(nums) && nums[child] > nums[child + 1] {
			child++
		}
		if nums[index] <= nums[child] {
			break
		}
		nums[index] = nums[child]
		//nums[index], nums[child] = nums[child], nums[index]
		index = child
		child = index << 1 + 1
	}
	nums[index] = temp

}

