package leetcode80

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}

	type result struct {
		k    int
		nums []int
	}

	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: result{
				k:    5,
				nums: []int{1, 1, 2, 2, 3},
			},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: result{
				k:    7,
				nums: []int{0, 0, 1, 1, 2, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)

			if got != tt.want.k {
				t.Errorf("removeDuplicates() = %v %v, want %v", got, tt.args.nums, tt.want)
				return
			}
			for i := 0; i < tt.want.k; i++ {
				if tt.args.nums[i] != tt.want.nums[i] {
					t.Errorf("removeDuplicates() = %v %v, want %v", got, tt.args.nums, tt.want)
					return
				}
			}
		})
	}
}
