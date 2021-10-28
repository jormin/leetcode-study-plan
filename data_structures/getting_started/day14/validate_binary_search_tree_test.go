package day14

import "testing"

// TestIsValidBST 测试验证二叉搜索树
func TestIsValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				root: &TreeNode{
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
			want: true,
		},
		{
			name: "02",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "03",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:  2,
						Left: nil,
					},
				},
			},
			want: false,
		},
		{
			name: "04",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isValidBST(tt.args.root); got != tt.want {
					t.Errorf("isValidBST() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
