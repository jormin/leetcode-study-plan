package day12

import "testing"

// TestClimbStairs 测试爬楼梯
func TestClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "02",
			args: args{n: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := climbStairs(tt.args.n); got != tt.want {
					t.Errorf("climbStairs() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
