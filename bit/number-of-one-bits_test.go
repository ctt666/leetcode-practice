package bit

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	count := 1
	tmp := num & (num - 1)
	for tmp != 0 {
		tmp = tmp & (tmp - 1)
		count++
	}
	return count
}
