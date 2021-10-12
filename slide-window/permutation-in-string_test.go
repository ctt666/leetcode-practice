package slide_window

//中心思想：如果s2包含字串等于s1的排列，则字母对应的数量相等
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	count := [26]int{}
	diff := 0
	for i, _ := range s1 {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for _, v := range count {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}

	for i := m; i < n; i++ {
		//进==出，不再计算
		left, right := s2[i-m]-'a', s2[i]-'a'
		if left == right {
			continue
		}

		if count[right] == 0 {
			diff++
		}
		//更新索引数量
		count[right]--
		//再次判断是否数量一致
		if count[right] == 0 {
			diff--
		}

		if count[left] == 0 {
			diff++
		}
		count[left]++
		if count[left] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}
	return false
}
