package leetcode

// func longestPalindrome(s string) string {
// 	start, maxLen := 0, 0
// 	expand := func(left, right int) {
// 		for left >= 0 && right < len(s) && s[left] == s[right] {
// 			left--
// 			right++
// 		}
// 		length := right - left - 1
// 		if length > maxLen {
// 			start = left + 1
// 			maxLen = length
// 		}
// 	}
// 	for i := 0; i < len(s); i++ {
// 		expand(i, i)
// 		expand(i, i+1)
// 	}
// 	return s[start : start+maxLen]

// }

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1
	for length := 1; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1

			if s[i] == s[j] && (length <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if length > maxLen {
					start, maxLen = i, length
				}
			}
		}
	}

	return s[start : start+maxLen]
}
