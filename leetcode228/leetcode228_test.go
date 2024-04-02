package leetcode228

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{nums: []int{0, 1, 2, 4, 5, 7}},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "example 2",
			args: args{nums: []int{0, 2, 3, 4, 6, 8, 9}},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "example 3",
			args: args{nums: []int{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
