package leetcode55

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "example 2",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "example 3",
			args: args{nums: []int{0, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
