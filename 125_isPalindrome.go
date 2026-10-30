package leetcode

import "strings"

func isAlphanumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		if left < right {
			cleft := strings.ToLower(string(s[left]))
			cright := strings.ToLower(string(s[right]))
			if cleft != cright {
				return false
			}
			left++
			right--
		}
	}
	return true
}
