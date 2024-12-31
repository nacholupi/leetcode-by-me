package stack

type aStack []rune

func (s *aStack) push(r rune) {
	*s = append(*s, r)
}

func (s *aStack) isEmpty() bool {
	return len(*s) == 0
}

func (s *aStack) pop() (rune, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, true
}

func IsValid(s string) bool {
	var st aStack

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
