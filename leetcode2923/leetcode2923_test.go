package leetcode2923

import "testing"

func Test_findChampion(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]int{
					{0, 1},
					{0, 0},
				},
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]int{
					{0, 0, 1},
					{1, 0, 1},
					{0, 0, 0},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findChampion(tt.args.grid); got != tt.want {
				t.Errorf("findChampion() = %v, want %v", got, tt.want)
			}
		})
	}
}
