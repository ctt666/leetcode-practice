package pointer

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
func sortedSquares(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	left, right := 0, length-1
	for i := length - 1; i >= 0; i-- {
		if nums[right]*nums[right] >= nums[left]*nums[left] {
			result[i] = nums[right] * nums[right]
			right--
		} else {
			result[i] = nums[left] * nums[left]
			left++
		}
	}
	return result
}
