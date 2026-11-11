package leetcode

import "strings"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result strings.Builder
	for index, value := range values {
		for num >= value {
			result.WriteString(symbols[index])
			num -= value
		}
	}
	return result.String()
}
