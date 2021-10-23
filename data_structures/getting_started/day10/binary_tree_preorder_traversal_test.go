package day10

import (
	"reflect"
	"testing"
)

// TestPreorderTraversal 测试二叉树的前序遍历
func TestPreorderTraversal(t *testing.T) {
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
			want: []int{1, 2, 3},
		},
		{
			name: "02",
			args: args{root: nil},
			want: nil,
		},
		{
			name: "03",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1},
		},
		{
			name: "04",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: []int{1, 2},
		},
		{
			name: "05",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: []int{1, 2},
		},
		{
			name: "06",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: []int{1, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
