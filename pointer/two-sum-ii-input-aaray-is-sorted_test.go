package pointer

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for right > left {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}
		left++
	}
	return []int{}
}
