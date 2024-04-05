package leetcode1026

import "testing"

func Test_maxAncestorDiff(t *testing.T) {
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
					Val: 8,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 1},
						Right: &TreeNode{
							Val:   6,
							Left:  &TreeNode{Val: 4},
							Right: &TreeNode{Val: 7},
						},
					},
					Right: &TreeNode{
						Val: 10,
						Right: &TreeNode{
							Val:  14,
							Left: &TreeNode{Val: 13},
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val:  0,
							Left: &TreeNode{Val: 3},
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAncestorDiff(tt.args.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
