package day13

import (
	"reflect"
	"testing"
)

// TestSearchBST 测试二叉搜索树中的搜索
func TestSearchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "01",
			args: args{
				root: tree,
				val:  2,
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "02",
			args: args{
				root: tree,
				val:  5,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("searchBST() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
