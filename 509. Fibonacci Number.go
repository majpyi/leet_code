package leet_code

func fib(N int) int {
	var backtrack func(N int) int
	backtrack = func(N int) int {
		if N == 1 {
			return 1
		}
		if N == 0 {
			return 0
		}
		return backtrack(N-1) + backtrack(N-2)
	}
	return backtrack(N)
}
