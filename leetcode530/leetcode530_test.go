package leetcode530

import "testing"

func Test_getMinimumDifference(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 6},
				},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   48,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 49},
					},
				},
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				root: &TreeNode{
					Val: 236,
					Left: &TreeNode{
						Val:   104,
						Right: &TreeNode{Val: 227},
					},
					Right: &TreeNode{
						Val:   701,
						Right: &TreeNode{Val: 911},
					},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
