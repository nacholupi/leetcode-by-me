package stack

//nolint:golint,unused
type stack []rune

//nolint:golint,unused
func (s *stack) push(r rune) {
	*s = append(*s, r)
}

//nolint:golint,unused
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

//nolint:golint,unused
func (s *stack) pop() (rune, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, true
}

//nolint:golint,unused, deadcode
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
