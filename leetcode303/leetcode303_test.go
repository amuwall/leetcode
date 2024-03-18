package leetcode303

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumArray(t *testing.T) {
	type arrayRange struct {
		left  int
		right int
	}

	type args struct {
		nums        []int
		arrayRanges []arrayRange
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:        []int{-2, 0, 3, -5, 2, -1},
				arrayRanges: []arrayRange{{0, 2}, {2, 5}, {0, 5}},
			},
			want: []int{1, -1, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.args.nums)
			for i, r := range tt.args.arrayRanges {
				if got := this.SumRange(r.left, r.right); got != tt.want[i] {
					t.Errorf("SumRange() = %v, want %v", got, tt.want[i])
					return
				}
			}
		})
	}
}
