package leetcode

// isValid use stack to check if the current close bracket closes its corresponding bracket
// problem: https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	parenthesesMapping := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	st := NewStack()
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			st = st.Push(c)
		case ')', ']', '}':
			var v int32
			st, v = st.Pop()
			if c != parenthesesMapping[v] {
				return false
			}
		default:
			return false
		}
	}
	return st.IsEmpty()
}

type Stack []int32

func NewStack(v ...int32) Stack {
	if v != nil {
		return v
	}
	return make([]int32, 0)
}

func (st Stack) Push(v int32) Stack {
	return append(st, v)
}

func (st Stack) Pop() (Stack, int32) {
	if st.IsEmpty() {
		return st, -1
	}
	l := len(st)
	return st[:l-1], st[l-1]
}

func (st Stack) IsEmpty() bool {
	return len(st) == 0
}
