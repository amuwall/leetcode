package leetcode189

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			wants: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			wants: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.wants) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.wants)
			}
		})
	}
}
