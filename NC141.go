package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge( str string ) bool {
	strs := []byte(str)

	for i := 0; i < len(strs) >> 1; i++ {
		if strs[i] != strs[len(strs) - 1 - i] {
			return false

		}
	}
	return true
}