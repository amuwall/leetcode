package leetcode98

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				root: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -2,
						Left: &TreeNode{
							Val: -3,
						},
					},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: true,
		},
		{
			name: "example 4",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
