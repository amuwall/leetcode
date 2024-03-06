package leetcode2917

import "testing"

func Test_findKOr(t *testing.T) {
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
				nums: []int{7, 12, 9, 8, 9, 15},
				k:    4,
			},
			want: 9,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 12, 1, 11, 4, 5},
				k:    6,
			},
			want: 0,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{10, 8, 5, 9, 11, 6, 8},
				k:    1,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKOr(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
