package day12

import "testing"

// TestHasPathSum 测试路径总和
func TestHasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
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
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   2,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:  4,
							Left: nil,
							Right: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "02",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "03",
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
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "04",
			args: args{
				root:      nil,
				targetSum: 1,
			},
			want: false,
		},
		{
			name: "05",
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
				targetSum: 1,
			},
			want: false,
		},
		{
			name: "06",
			args: args{
				root: &TreeNode{
					Val: -2,
					Left: nil,
					Right: &TreeNode{
						Val:   -3,
						Left:  nil,
						Right: nil,
					},
				},
				targetSum: -5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
					t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
