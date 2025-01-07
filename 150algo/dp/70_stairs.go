package dp

func climbStairs(n int) int {
	return climbStairsMap(n, make(map[int]int))
}

func climbStairsMap(n int, m map[int]int) int {
	if n < 3 {
		return n
	}

	v, ok := m[n]
	if ok {
		return v
	}

	result := climbStairsMap(n-2, m) + climbStairsMap(n-1, m)
	m[n] = result
	return result
}

/*
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    res, first, second := 0,2,3
    for i := 0; i < n-3; i++ {
        res = first+second
        first, second = second, res
    }
    return res
}
*/
