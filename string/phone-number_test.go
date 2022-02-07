package string

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var result []string

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result = []string{}
	backTrack(digits, 0, "")
	return result
}

func backTrack(digits string, index int, combination string) {
	if index == len(digits) {
		result = append(result, combination)
		return
	}
	digit := string(digits[index])
	for _, v := range m[digit] {
		backTrack(digits, index+1, combination+string(v))
	}
}
