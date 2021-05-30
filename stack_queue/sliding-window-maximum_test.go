package stack_queue

func maxSlidingWindow(nums []int, k int) []int {
	//传入参数校验
	if nums == nil || len(nums) == 0 {
		return nil
	}
	var result []int
	var window []int
	for i := range nums {
		//判断是否pop
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		//判断插入前是否需要迁移之前的数据
		for len(window) > 0 && nums[i] > nums[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		//判断是否写入，并写入结果
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
