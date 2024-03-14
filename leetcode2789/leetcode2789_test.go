package leetcode2789

import "testing"

func Test_maxArrayValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{2, 3, 7, 9, 3},
			},
			want: 21,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{5, 3, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArrayValue(tt.args.nums); got != tt.want {
				t.Errorf("maxArrayValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
