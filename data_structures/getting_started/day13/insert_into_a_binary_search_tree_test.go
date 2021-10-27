package day13

import (
	"reflect"
	"testing"
)

// TestInsertIntoBST 测试二叉搜索树中的插入操作
func TestInsertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "01",
			args: args{
				root: &TreeNode{
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
				},
				val: 5,
			},
			want: &TreeNode{
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
					Val: 7,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			name: "02",
			args: args{
				root: &TreeNode{
					Val: 40,
					Left: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val:   10,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   30,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 60,
						Left: &TreeNode{
							Val:   50,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   70,
							Left:  nil,
							Right: nil,
						},
					},
				},
				val: 25,
			},
			want: &TreeNode{
				Val: 40,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   10,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 30,
						Left: &TreeNode{
							Val:   25,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 60,
					Left: &TreeNode{
						Val:   50,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   70,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "03",
			args: args{
				root: &TreeNode{
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
				},
				val: 8,
			},
			want: &TreeNode{
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
					Val:  7,
					Left: nil,
					Right: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "04",
			args: args{
				root: nil,
				val:  1,
			},
			want: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
