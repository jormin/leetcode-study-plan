package day14

import (
	"reflect"
	"testing"
)

// TestLowestCommonAncestor 测试二叉搜索树的最近公共祖先
func TestLowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	node1 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	node3 := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
	}
	node4 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	node2.Right = node4
	node1.Left = node2
	node1.Right = node3

	node5 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node7 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node8 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node5.Right = node6
	node6.Left = node7
	node6.Right = node8
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "01",
			args: args{
				root: node1,
				p:    node2,
				q:    node3,
			},
			want: node1,
		},
		{
			name: "02",
			args: args{
				root: node1,
				p:    node2,
				q:    node4,
			},
			want: node2,
		},
		{
			name: "03",
			args: args{
				root: node1,
				p:    node7,
				q:    node8,
			},
			want: node6,
		},
		{
			name: "04",
			args: args{
				root: nil,
				p:    node7,
				q:    node8,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
