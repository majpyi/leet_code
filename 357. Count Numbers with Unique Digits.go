package leet_code

import (
	"fmt"
)

func countNumbersWithUniqueDigits(n int) int {
	re := 0
	arr := make([]int, 0, n)
	traceBack(n, arr, &re)
	return re
}

func traceBack(n int, arr []int, re *int) {
	if n == len(arr) {
		fmt.Println(arr)
		*re++
		return
	}
	for i := 0; i <= 9; i++ {
		if isNotIn(i, arr) {
			arr = append(arr, i)
			traceBack(n, arr, re)
			arr = arr[:len(arr)-1]
		}
	}
}

func isNotIn(i int, arr []int) bool {
	j := 0
	for ; j < len(arr); j++ {
		if arr[j] != 0 {
			break
		}
		if j == len(arr)-1 {
			return true
		}
	}
	for x := j; x < len(arr); x++ {
		if arr[x] == i {
			return false
		}
	}
	return true
}
