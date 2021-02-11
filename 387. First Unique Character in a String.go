package leet_code

func firstUniqChar(s string) int {
	m:=make(map[int32]int,len(s))
	for index,char :=range s{
		if _,ok:=m[char-'a'];ok{
			m[char-'a'] = -1
		}else{
			m[char-'a'] = index
		}
	}
	for index,char :=range s{
		if m[char-'a']>0{
			return index
		}
	}
	return -1
}
