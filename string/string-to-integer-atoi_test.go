package string

import (
	"math"
)

func myAtoi(s string) (ans int) {
	table := map[string][]string{
		"start":     []string{"start", "signed", "in_number", "end"},
		"signed":    []string{"end", "end", "in_number", "end"},
		"in_number": []string{"end", "end", "in_number", "end"},
	}
	state := "start"
	sign := 1
	for i, _ := range s {
		if state == "end" {
			break
		}
		if s[i] == ' ' {
			state = table[state][0]
			continue
		}
		if s[i] == '+' || s[i] == '-' {
			state = table[state][1]
			if s[i] == '-' && state == "signed" {
				sign = -1
			}
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			state = table[state][2]
			if sign == 1 && (math.MaxInt32-int(s[i]-'0'))/10 < ans {
				ans = math.MaxInt32
				return
			}
			if sign == -1 && -(math.MinInt32+int(s[i]-'0'))/10 < ans {
				ans = math.MinInt32
				return
			}
			ans = ans*10 + int(s[i]-'0')
		} else {
			state = table[state][3]
		}
	}

	ans = sign * ans
	return
}
