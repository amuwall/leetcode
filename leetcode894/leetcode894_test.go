package leetcode894

import (
	"reflect"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "example 1",
			args: args{n: 7},
			want: []*TreeNode{
				{
					Left: &TreeNode{},
					Right: &TreeNode{
						Left: &TreeNode{},
						Right: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
					},
				},
				{
					Left: &TreeNode{},
					Right: &TreeNode{
						Left: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
						Right: &TreeNode{},
					},
				},
				{
					Left: &TreeNode{
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
					Right: &TreeNode{
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
				},
				{
					Left: &TreeNode{
						Left: &TreeNode{},
						Right: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
					},
					Right: &TreeNode{},
				},
				{
					Left: &TreeNode{
						Left: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
						Right: &TreeNode{},
					},
					Right: &TreeNode{},
				},
			},
		},
		{
			name: "example 2",
			args: args{n: 3},
			want: []*TreeNode{
				{
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allPossibleFBT(tt.args.n)
			if len(got) != len(tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
				return
			}
			for i := 0; i < len(got); i++ {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
