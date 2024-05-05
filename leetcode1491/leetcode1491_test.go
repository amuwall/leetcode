package leetcode1491

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example 1",
			args: args{
				salary: []int{4000, 3000, 1000, 2000},
			},
			want: 2500,
		},
		{
			name: "example 2",
			args: args{
				salary: []int{1000, 2000, 3000},
			},
			want: 2000,
		},
		{
			name: "example 3",
			args: args{
				salary: []int{6000, 5000, 4000, 3000, 2000, 1000},
			},
			want: 3500,
		},
		{
			name: "example 4",
			args: args{
				salary: []int{8000, 9000, 2000, 3000, 6000, 1000},
			},
			want: 4750,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
