package leetcode2529

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{-2, -1, -1, 1, 2, 3}},
			want: 3,
		},
		{
			name: "example 2",
			args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}},
			want: 3,
		},
		{
			name: "example 3",
			args: args{nums: []int{5, 20, 66, 1314}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
