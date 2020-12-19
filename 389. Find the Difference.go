package leet_code

import (
	"strings"
)

func findTheDifference(s string, t string) byte {
	for _, c := range s {
		for i, d := range t {
			if c == d {
				//t = t[:i]+[i+1:]
				//t = t[:i]
				t = strings.Join([]string{t[:i], t[i+1:]}, "")
				break
			}
		}
	}
	return []byte(t)[0]
}

//func findTheDifference(s, t string) byte {
//	cnt := [26]int{}
//	for _, ch := range s {
//		cnt[ch-'a']++
//	}
//	for i := 0; ; i++ {
//		ch := t[i]
//		cnt[ch-'a']--
//		if cnt[ch-'a'] < 0 {
//			return ch
//		}
//	}
//}
