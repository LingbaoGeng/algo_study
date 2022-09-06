package dp

import "algo_study/util"

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = util.Min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
