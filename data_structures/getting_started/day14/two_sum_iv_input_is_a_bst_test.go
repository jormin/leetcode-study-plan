package day14

import "testing"

// TestFindTarget 测试两数之和 IV - 输入 BST
func TestFindTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				root: nil,
				k:    9,
			},
			want: false,
		},
		{
			name: "02",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
				k: 9,
			},
			want: true,
		},
		{
			name: "03",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
				k: 28,
			},
			want: false,
		},
		{
			name: "04",
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
				k: 4,
			},
			want: true,
		},
		{
			name: "05",
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
				k: 1,
			},
			want: false,
		},
		{
			name: "06",
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
				k: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
					t.Errorf("findTarget() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
