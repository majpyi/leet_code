package leet_code

func generate(numRows int) [][]int {
	re := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		re[i] = make([]int, 0, i+1)
		re[i][0] = 1
		re[i][i] = 1
		for j := 1; j < i; j++ {
			re[i][j] = re[i-1][j-1] + re[i-1][j]
		}

	}
	return re
}

func generateBest(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
