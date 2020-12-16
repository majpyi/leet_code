package leet_code

func generateParenthesis(n int) []string {
	var re = &[]string{}
	traceBack1("", n, 0, 0, n, n, re)
	return *re
}

func traceBack1(s string, n, l, r, ll, rr int, re *[]string) {
	if len(s) == n*2 {
		*re = append(*re, s)
		return
	}
	if l >= r && ll > 0 {
		traceBack1(s+"(", n, l+1, r, ll-1, rr, re)
	}
	if l >= r && rr > 0 {
		traceBack1(s+")", n, l, r+1, ll, rr-1, re)
	}
	return
}
