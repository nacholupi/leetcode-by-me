package array

func maxProfit(prices []int) int {
	grtst := 0
	if len(prices) > 1 {
		for i := 0; i < len(prices); i++ {
			for j := i + 1; j < len(prices); j++ {
				if prices[i] < prices[j] {
					grtst = max(grtst, prices[j]-prices[i])
				} else {
					break
				}
			}
		}
	}
	return grtst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
