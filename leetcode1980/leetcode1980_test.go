package leetcode1980

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				nums: []string{"01", "10"},
			},
			want: "11",
		},
		{
			name: "example 2",
			args: args{
				nums: []string{"00", "01"},
			},
			want: "10",
		},
		{
			name: "example 3",
			args: args{
				nums: []string{"111", "011", "001"},
			},
			want: "000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.args.nums); got != tt.want {
				t.Errorf("findDifferentBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
