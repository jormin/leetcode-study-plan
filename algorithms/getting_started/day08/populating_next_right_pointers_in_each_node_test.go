package day08

import (
	"reflect"
	"testing"
)

// TestConnect 测试填充每个节点的下一个右侧节点指针
func TestConnect(t *testing.T) {
	type args struct {
		root *Node
	}
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node1.Left = node2
	node1.Next = nil
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node2.Next = node3
	node3.Left = node6
	node3.Right = node7
	node3.Next = nil
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = nil

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "01",
			args: args{
				root: &Node{
					Val: 1,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &Node{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &Node{
						Val: 3,
						Left: &Node{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &Node{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: node1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("connect() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
