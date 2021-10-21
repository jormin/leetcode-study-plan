package day07

import "testing"

// TestHasCycle 测试环形链表
func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node1 := &ListNode{
		Val: 3,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 0,
	}
	node4 := &ListNode{
		Val: -4,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	node5 := &ListNode{
		Val: 1,
	}
	node6 := &ListNode{
		Val: 2,
	}
	node5.Next = node6
	node6.Next = node5
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{head: nil},
			want: false,
		},
		{
			name: "02",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: false,
		},
		{
			name: "03",
			args: args{head: node1},
			want: true,
		},
		{
			name: "04",
			args: args{head: node5},
			want: true,
		},
		{
			name: "05",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := hasCycle(tt.args.head); got != tt.want {
					t.Errorf("hasCycle() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
