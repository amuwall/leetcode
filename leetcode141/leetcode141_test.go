package leetcode141

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  -4,
								Next: nil,
							},
						},
					},
				},
				pos: 1,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				pos: 0,
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				head: &ListNode{Val: 1},
				pos:  -1,
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{head: nil, pos: -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.pos >= 0 {
				p := tt.args.head
				var q *ListNode
				var i int
				for p.Next != nil {
					if i == tt.args.pos {
						q = p
					}

					i++
					p = p.Next
				}
				p.Next = q
			}

			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
