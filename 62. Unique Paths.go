package leet_code

func uniquePaths(m int, n int) int {
	arr := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		arr[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		arr[0][i] = 1
	}
	for i := 1; i <= m; i++ {
		arr[i][0] = 1
		for j := 1; j <= n; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
	}
	return arr[m-1][n-1]
}
