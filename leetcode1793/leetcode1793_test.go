package leetcode1793

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 4, 3, 7, 4, 5},
				k:    3,
			},
			want: 15,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{5, 5, 4, 5, 4, 1, 1, 1},
				k:    0,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
