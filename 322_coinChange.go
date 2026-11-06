package leetcode

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i-coin >= 0 && (dp[i-coin] != -1 || i-coin == 0) {
				if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}
