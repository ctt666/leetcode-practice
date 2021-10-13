package pointer

func moveZeroes(nums []int) {
	left, right := 0, 0
	n := len(nums)
	for right < n && left < n {
		if nums[right] == 0 {
			right++
			continue
		}
		if nums[left] != 0 {
			left++
			continue
		}
		if right > left {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
}

func moveZeroes_01(nums []int) {
	left, right := 0, 0
	n := len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
