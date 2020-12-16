package leet_code

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	arrStr := strings.Split(s, " ")
	if len(pattern) != len(arrStr) {
		return false
	}

	p := make(map[int32]*[]int, 26)
	for i, c := range pattern {
		if pp, ok := p[c]; ok {
			*pp = append(*pp, i)
		} else {
			p1 := &[]int{i}
			p[c] = p1
		}
	}
	f := make(map[string]bool, 26)
	for _, v := range p {
		t := ""
		if f[arrStr[(*v)[0]]] {
			return false
		}
		f[arrStr[(*v)[0]]] = true
		for _, i := range *v {
			if arrStr[i] != t && t != "" {
				return false
			}
			t = arrStr[i]
		}
	}
	return true
}
