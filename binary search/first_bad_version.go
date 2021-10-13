package binary_search

func firstBadVersion(n int) int {
	oldest, latest := 1, n
	for oldest <= latest {
		mid := (oldest + latest) / 2
		if isBadVersion(mid) && (mid == 1 || !isBadVersion(mid-1)) {
			return mid
		}
		if isBadVersion(mid) {
			latest = mid - 1
		} else {
			oldest = mid + 1
		}
	}
	return -1
}

//mock
func isBadVersion(version int) bool {
	return true
}
