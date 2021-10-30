package day09

import "testing"

// TestOrangesRotting 测试腐烂的橘子
func TestOrangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			want: 4,
		},
		{
			name: "02",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			want: -1,
		},
		{
			name: "02",
			args: args{
				grid: [][]int{
					{0, 2},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := orangesRotting(tt.args.grid); got != tt.want {
					t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
