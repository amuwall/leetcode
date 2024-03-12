package leetcode1261

import (
	"testing"
)

func TestFindElements_Find(t *testing.T) {
	type args struct {
		root    *TreeNode
		targets []int
	}
	tests := []struct {
		name  string
		args  args
		wants []bool
	}{
		{
			name: "example 1",
			args: args{
				root: &TreeNode{
					Val:  -1,
					Left: nil,
					Right: &TreeNode{
						Val: -1,
					},
				},
				targets: []int{1, 2},
			},
			wants: []bool{false, true},
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -1,
						Left: &TreeNode{
							Val:   -1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   -1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: -1,
					},
				},
				targets: []int{1, 3, 5},
			},
			wants: []bool{true, true, false},
		},
		{
			name: "example 3",
			args: args{
				root: &TreeNode{
					Val:  -1,
					Left: nil,
					Right: &TreeNode{
						Val: -1,
						Left: &TreeNode{
							Val: -1,
							Left: &TreeNode{
								Val:   -1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
				},
				targets: []int{2, 3, 4, 5},
			},
			wants: []bool{true, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.args.root)
			for i, target := range tt.args.targets {
				if got := this.Find(target); got != tt.wants[i] {
					t.Errorf("Find(%d) = %v, want %v", target, got, tt.wants[i])
					return
				}
			}
		})
	}
}
