package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPalindrome(tt.args.x), "isPalindrome(%v)", tt.args.x)
		})
	}
}
