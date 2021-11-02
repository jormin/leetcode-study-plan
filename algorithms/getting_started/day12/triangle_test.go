package day12

import "testing"

// TestMinimumTotal 测试三角形最小路径和
func TestMinimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				triangle: [][]int{
					{2},
					{3, 4},
					{6, 5, 7},
					{4, 1, 8, 3},
				},
			},
			want: 11,
		},
		{
			name: "02",
			args: args{
				triangle: [][]int{
					{-10},
				},
			},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minimumTotal(tt.args.triangle); got != tt.want {
					t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
