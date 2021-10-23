package day10

import (
	"reflect"
	"testing"
)

// 测试二叉树的后序遍历
func TestPostorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
