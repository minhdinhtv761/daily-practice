package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				bills: []int{5, 5, 5, 10, 20},
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				bills: []int{5, 5, 10, 10, 20},
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				bills: []int{5, 5, 5, 5, 20, 20, 5, 5, 5, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lemonadeChange(tt.args.bills), "lemonadeChange(%v)", tt.args.bills)
		})
	}
}
