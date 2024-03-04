package leetcode21

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
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
									Val:  4,
									Next: nil,
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
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "example 3",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)

			p := got
			q := tt.want
			for {
				if p == nil && q == nil {
					break
				}

				if (p == nil && q != nil) || (p != nil && q == nil) {
					t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
					return
				}
				if p.Val != q.Val {
					t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
					return
				}

				p = p.Next
				q = q.Next
			}
		})
	}
}
