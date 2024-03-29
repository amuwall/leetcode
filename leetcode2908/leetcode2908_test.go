package leetcode2908

import "testing"

func Test_minimumSum(t *testing.T) {
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
			args: args{
				nums: []int{8, 6, 1, 5, 3},
			},
			want: 9,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{5, 4, 8, 7, 10, 2},
			},
			want: 13,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{6, 5, 4, 3, 4, 5},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.nums); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
