package leetcode129

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 25,
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   9,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 0},
				},
			},
			want: 1026,
		},
		{
			name: "example 3",
			args: args{
				root: &TreeNode{Val: 9},
			},
			want: 9,
		},
		{
			name: "example 4",
			args: args{
				root: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 1},
				},
			},
			want: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
