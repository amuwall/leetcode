package leetcode23

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				lists: []*ListNode{
					{
						Val: 1,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5},
						},
					},
					{
						Val: 1,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{Val: 4},
						},
					},
					{
						Val:  2,
						Next: &ListNode{Val: 6},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  5,
										Next: &ListNode{Val: 6},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "example 2",
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			name: "example 3",
			args: args{
				lists: []*ListNode{nil},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
