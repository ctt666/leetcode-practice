package dynamic

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	f1 := nums[0]
	f2 := max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		f1, f2 = f2, max(nums[i]+f1, f2)
	}
	return f2
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
