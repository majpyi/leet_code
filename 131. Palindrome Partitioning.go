package leet_code

func checkP(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func recurseP(cur []string, ret *[][]string, index int, s string) {
	if index == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}

	for i := index; i < len(s); i++ {
		if checkP(s[index : i+1]) {
			cur = append(cur, s[index:i+1])
			recurseP(cur, ret, i+1, s)
			cur = cur[:len(cur)-1]
		}
	}
}

func partition(s string) [][]string {
	ret := make([][]string, 0)
	cur := make([]string, 0)
	recurseP(cur, &ret, 0, s)
	return ret
}
