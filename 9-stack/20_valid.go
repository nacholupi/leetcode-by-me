package stack

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func rotate(nums []int, k int) {
	left := nums[:k]
	right := nums[k+1:]

	copy(nums, right)

	for j := range left {
		nums[j+k] = left[j]
	}
}

func (s *stack) pop() (rune, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, true
}

func isValid(s string) bool {
	var st stack

	for _, c := range s {
		switch c {
		case '(':
			st.push(')')
		case '[':
			st.push(']')
		case '{':
			st.push('}')
		default:
			e, ok := st.pop()
			if !ok || e != c {
				return false
			}
		}
	}
	return st.isEmpty()
}
