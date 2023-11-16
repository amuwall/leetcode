package leetcode1535

import "testing"

func Test_getWinner(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{2, 1, 3, 5, 4, 6, 7},
				k:   2,
			},
			want: 5,
		},
		{
			name: "example 2",
			args: args{
				arr: []int{3, 2, 1},
				k:   10,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinner(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
