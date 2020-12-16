package leet_code

import (
	"strings"
)

func sortString(s string) string {
	var hash [26]byte
	var b strings.Builder
	for _, c := range s {
		hash[c-'a']++
	}
	rest := len(s)
	for rest > 0 {
		for i := 0; i < 26; i++ {
			if hash[i] > 0 {
				b.WriteByte('a' + byte(i))
				hash[i]--
				rest--
			}
		}
		for i := 25; i >= 0; i-- {
			if hash[i] > 0 {
				b.WriteByte('a' + byte(i))
				hash[i]--
				rest--
			}
		}
	}
	return b.String()
}
