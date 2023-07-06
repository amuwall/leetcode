package leetcode2

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "example 2",
			args: args{
				l1: &ListNode{
					Val:  0,
					Next: nil,
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "example 3",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			want := tt.want
			for {
				if want.Val != got.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got, want)
					break
				}

				want = want.Next
				got = got.Next

				if (want != nil && got == nil) || (want == nil && got != nil) {
					t.Errorf("addTwoNumbers() = %v, want %v", got, want)
					break
				}

				if want == nil && got == nil {
					break
				}
			}
		})
	}
}
