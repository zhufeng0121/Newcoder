package main

import "fmt"

//最长公共子序列

/**
 * longest common subsequence
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
 */

//此题思路很明显是DP，但是DP数组的构造却是一定要慎之又慎  dp [m + 1][n + 1]
func LCS( s1 string ,  s2 string ) int {
	m, n  := len(s1), len(s2)
	dp := make([][]int, m + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	//dp[i][j] 以 nums1[i]和nums2[j]为结尾的两个数组中公共的、长度最长的子数组的长度
	//res := make([]byte,0)
	//dp[m][n]代表什么含义呢,即长度为m和n的字符串公共子串的长度
	//dp[0][0] = 0 dp[m][0] = 0 dp[0][n] = 0
	for i := 1 ; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = maxI(dp[i-1][j],dp[i][j-1])
		}
	}
	return dp[m][n]

}

func maxI(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	s1 := "1A2C3D4B56"
	s2 := "B1D23A456A"
	fmt.Println(LCS(s1,s2))
}
