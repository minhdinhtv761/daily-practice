package models

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
