package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name                 string
		args                 args
		want                 int
		wantOrderedUniqeNums []int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:                 2,
			wantOrderedUniqeNums: []int{1, 2},
		},
		{
			name: "test case 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:                 5,
			wantOrderedUniqeNums: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.args.nums)
			assert.Equalf(t, tt.want, result, "removeDuplicates(%v)", tt.args.nums)
			assert.Equal(t, tt.wantOrderedUniqeNums, tt.args.nums[:result])
		})
	}
}
