package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find( target int ,  array [][]int ) bool {
	row := 0
	column := len(array[0]) - 1

	for row < len(array) && column >= 0 {
		if target == array[row][column] {
			return true
		} else if target < array[row][column] {
			column--
		} else if target > array[row][column] {
			row++
		}
	}
	return false
}
