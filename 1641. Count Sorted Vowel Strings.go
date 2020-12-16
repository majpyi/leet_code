package leet_code

func countVowelStrings(n int) int {
	dp := make([][5]int, n+1)
	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		//长度i的以u结尾的字符串可以由任意一个长度i-1的字符串结尾加个u得到
		dp[i][4] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + dp[i-1][4]
		dp[i][3] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3]
		dp[i][2] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2]
		dp[i][1] = dp[i-1][0] + dp[i-1][1]
		//长度i的以a结尾的字符串只能由长度i-1的以a结尾的字符串结尾加个a得到
		dp[i][0] = dp[i-1][0]
	}
	return dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]
}

func countVowelStrings1(n int) int {
	dp := make([][5]int, n+1)
	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i][0] + dp[i-1][1]
		dp[i][2] = dp[i][1] + dp[i-1][2]
		dp[i][3] = dp[i][2] + dp[i-1][3]
		dp[i][4] = dp[i][3] + dp[i-1][4]
	}
	return dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]
}
