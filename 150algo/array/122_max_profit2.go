package array

// Consider this input : [a, ......, b, ......., c....., d]. Where ... refers to some input in between. a&c are valleys and b&d are peaks.
// If b&d both are peaks then there has to be some value in between them which is less than b. If not b could not have been a peak because the direct next peak immediate to a will be d.
// If a&c are valleys then there has to be some value in between them which is greater than a. If not a could not have been a valley because the first valley will itself will be c.
// we can always say that picking up immediate next peak will always yeild a better result.
func maxProfit2(prices []int) int {
	grtst := 0
	for i := 0; i < len(prices)-1; i++ {

		if prices[i+1] > prices[i] {
			grtst = grtst + prices[i+1] - prices[i]
		}
	}
	return grtst
}
