package leetcode415

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		},
		{
			name: "example 2",
			args: args{
				num1: "456",
				num2: "77",
			},
			want: "533",
		},
		{
			name: "example 3",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "example 4",
			args: args{
				num1: "99",
				num2: "1",
			},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
