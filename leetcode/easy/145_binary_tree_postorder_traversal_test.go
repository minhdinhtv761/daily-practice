package easy

import (
	"daily-practice/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *models.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				root: &models.TreeNode{
					Val:  1,
					Left: nil,
					Right: &models.TreeNode{
						Val: 2,
						Left: &models.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "test case 2",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "test case 3",
			args: args{
				root: &models.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, postorderTraversal(tt.args.root), "postorderTraversal(%v)", tt.args.root)
		})
	}
}
