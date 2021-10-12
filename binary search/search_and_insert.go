package binary_search

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := len(nums)
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] >= target {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
