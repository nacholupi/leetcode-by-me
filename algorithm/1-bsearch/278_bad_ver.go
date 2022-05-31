package bsearch

func firstBadVersion(n int) int {
	start := 1
	end := n
	bad := 0

	for end >= start {
		h := (start + end) / 2
		if isBadVersion(h) {
			bad = h
			end = h - 1
		} else {
			start = h + 1
		}
	}
	return bad
}

func isBadVersion(version int) bool {
	return version >= 2
}
