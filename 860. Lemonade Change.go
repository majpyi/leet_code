package leet_code

func lemonadeChange(bills []int) bool {
	arr := [3]int{}
	for _, bill := range bills {
		if bill == 5 {
			arr[0]++
		}
		if bill == 10 {
			arr[0]--
			if arr[0] < 0 {
				return false
			}
			arr[1]++
		}
		if bill == 20 {
			if arr[1] > 0 {
				arr[1]--
				arr[0]--
				if arr[0] < 0 {
					return false
				}
			} else {
				arr[0] -= 3
				if arr[0] < 0 {
					return false
				}
			}
		}
	}
	return true
}
