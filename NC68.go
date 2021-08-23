package main

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor( number int ) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	a, b := 1, 2
	res := 0
	for i := 3; i <= number; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}