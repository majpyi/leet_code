package leet_code

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[int32]int64, len(s))

	for _, c := range s {
		if _, ok := charMap[c]; ok {
			charMap[c] += 1
		} else {
			charMap[c] = 1
		}
	}
	for _, c := range t {
		charMap[c] -= 1
		if charMap[c] < 0 {
			return false
		}
	}
	return true
}
