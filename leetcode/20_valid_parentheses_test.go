package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "test case 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValid(tt.args.s), "isValid(%v)", tt.args.s)
		})
	}
}
