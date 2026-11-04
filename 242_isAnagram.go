package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[rune]int)
	for _, c := range s {
		sm[c]++
	}

	for _, c := range t {
		if sm[c] == 0 {
			return false
		}
		sm[c]--
	}
	return true
}
