package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "test case ",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, romanToInt(tt.args.s), "romanToInt(%v)", tt.args.s)
		})
	}
}
