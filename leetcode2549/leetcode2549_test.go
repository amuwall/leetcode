package leetcode2549

import "testing"

func Test_distinctIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{n: 5},
			want: 4,
		},
		{
			name: "example 2",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "example 3",
			args: args{n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctIntegers(tt.args.n); got != tt.want {
				t.Errorf("distinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
