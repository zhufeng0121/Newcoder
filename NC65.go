package main


func Fibonacci( n int ) int {
	// write code here
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
