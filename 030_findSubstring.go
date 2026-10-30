package leetcode

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 || len(words[0]) == 0 {
		return res
	}
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	if totalLen > len(s) {
		return res
	}
	need := make(map[string]int, len(words))
	for _, word := range words {
		need[word]++
	}
	for offset := 0; offset < wordLen && offset+totalLen <= len(s); offset++ {
		window := make(map[string]int, len(need))
		left, count := offset, 0
		for right := offset; right+wordLen <= len(s); right += wordLen {
			word := s[right : right+wordLen]
			if need[word] == 0 {
				clear(window)
				left, count = right+wordLen, 0
				continue
			}
			window[word]++
			count++
			for window[word] > need[word] {
				first := s[left : left+wordLen]
				window[first]--
				count--
				left += wordLen
			}
			if count == len(words) {
				res = append(res, left)
				first := s[left : left+wordLen]
				window[first]--
				count--
				left += wordLen
			}
		}
	}
	return res
}
