package leet_code

import "strconv"

func isAdditiveNumber(s string) bool {
	if len(s) < 3 { // 长度不够
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[0] == '0' && i != 0 { // 第一个数字有前导 0
			break
		}
		first, _ := strconv.Atoi(s[:i+1])
		for j := i + 1; j < len(s); j++ { // 第二个数字有前导 0
			if s[i+1] == '0' && j != i+1 {
				break
			}
			second, _ := strconv.Atoi(s[i+1 : j+1])
			if checkNext(s[j+1:], first, second, 2) { // 分解成功则返回 true
				return true
			}
		}
	}
	return false
}

func checkNext(s string, first, second int, count int) bool {
	if len(s) == 0 { // 分解完成之后判断数字个数是否大于等于 3
		return count >= 3
	}
	num := 0
	for i, c := range s {
		if s[0] == '0' && i != 0 { // 有前导 0
			break
		}
		num = num*10 + int(c-'0')
		if num > first+second { // 大于前两数之和，返回 false
			return false
		} else if num == first+second { // 递归向后判断
			return checkNext(s[i+1:], second, num, count+1)
		}
	}
	return false
}
