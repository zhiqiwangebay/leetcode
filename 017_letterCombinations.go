package leetcode

func letterCombinations(digits string) []string {
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	if len(digits) == 0 {
		return []string{}
	}
	res := []string{}
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := digitMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, path+string(letters[i]))
		}
	}
	backtrack(0, "")
	return res
}
