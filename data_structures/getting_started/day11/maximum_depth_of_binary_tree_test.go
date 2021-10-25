package day11

import "testing"

// TestMaxDepth 测试二叉树的最大深度
func TestMaxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{root: nil},
			want: 0,
		},
		{
			name: "02",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val:   15,
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := maxDepth(tt.args.root); got != tt.want {
					t.Errorf("maxDepth() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
