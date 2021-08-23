package main

import "strings"

//大数加法
func solve(s string, t string) string {
    if len(s) > len(t) {
    	s, t = t, s
	}
	s = strings.Repeat("0",len(t) - len(s)) + s

	strs := []byte(s)
	strt := []byte(t)
	carry := byte(0)
	res := make([]byte,len(strs))
	for i := len(strs) - 1; i >= 0; i-- {
		//这里的字符转化一定要记得加上字符0！！！！！！！
		res[i] = (carry + strs[i] - '0' + strt[i] - '0') % 10 + '0'
		carry  = (carry + strs[i] - '0' + strt[i] - '0') / 10
	}
	if carry == 1 {
		return "1" + string(res)
	}
	return string(res)
}
