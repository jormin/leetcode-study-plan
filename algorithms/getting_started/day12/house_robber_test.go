package day12

import "testing"

// TestRob 测试打家劫舍
func TestRob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "02",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "03",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "04",
			args: args{nums: []int{2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := rob(tt.args.nums); got != tt.want {
					t.Errorf("rob() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
