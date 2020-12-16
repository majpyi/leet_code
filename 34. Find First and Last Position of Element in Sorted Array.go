package leet_code

func searchRange(nums []int, target int) []int {
	r := len(nums) - 1
	l := 0
	result_l := -1
	result_r := -1
	for {
		if l > r {
			return []int{result_l, result_r}
		}
		mid := l + (r-l)/2
		if nums[mid] == target {
			result_l = mid
			result_r = mid
			for i := mid - 1; i >= 0; i-- {
				if nums[i] == target {
					result_l = i
					continue
				}
				break
			}
			for i := mid; i < len(nums); i++ {
				if nums[i] == target {
					result_r = i
					continue
				}
				break
			}
			return []int{result_l, result_r}

		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
