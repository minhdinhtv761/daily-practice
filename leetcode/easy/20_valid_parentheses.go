package easy

import "daily-practice/models"

// isValid use stack to check if the current close bracket closes its corresponding bracket
// problem: https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	parenthesesMapping := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	st := models.NewStack()
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
