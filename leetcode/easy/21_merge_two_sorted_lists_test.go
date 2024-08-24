package easy

import (
	"daily-practice/leetcode/global"
	"daily-practice/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *models.ListNode
		list2 *models.ListNode
	}
	tests := []struct {
		name string
		args args
		mode string
		want []int
	}{
		{
			name: "recursive - test case 1",
			args: args{
				list1: models.NewListNode([]int{1, 2, 4}),
				list2: models.NewListNode([]int{1, 3, 4}),
			},
			mode: global.SolutionModeRecursive,
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "recursive - test case 2",
			args: args{
				list1: models.NewListNode([]int{}),
				list2: models.NewListNode([]int{}),
			},
			mode: global.SolutionModeRecursive,
			want: []int{},
		},
		{
			name: "recursive - test case 3",
			args: args{
				list1: models.NewListNode([]int{}),
				list2: models.NewListNode([]int{0}),
			},
			mode: global.SolutionModeRecursive,
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			global.SetSolutionMode(tt.mode)
			assert.Equalf(t, tt.want, mergeTwoLists(tt.args.list1, tt.args.list2).GetVals(), "mergeTwoLists(%v, %v)", tt.args.list1, tt.args.list2)
		})
	}
}
