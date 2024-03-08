package leetcode2834

import "testing"

func Test_minimumPossibleSum(t *testing.T) {
	type args struct {
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:      2,
				target: 3,
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				n:      3,
				target: 3,
			},
			want: 8,
		},
		{
			name: "example 3",
			args: args{
				n:      1,
				target: 1,
			},
			want: 1,
		},
		{
			name: "example 4",
			args: args{
				n:      39636,
				target: 49035,
			},
			want: 156198582,
		},
		{
			name: "example 5",
			args: args{
				n:      1,
				target: 1000000000,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPossibleSum(tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minimumPossibleSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
