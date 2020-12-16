package leet_code

func countPrimes1(n int) int {
	re := 0
	arr := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if i*j > n {
				break
			}
			arr[i*j]++
		}
	}
	for i := 2; i < n; i++ {
		if arr[i] == 0 {
			re++
		}
	}
	return re
}

func countPrimes2(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
