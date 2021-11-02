package day05

import "testing"

// TestSubarraySum 测试和为 K 的子数组
func TestSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "02",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
					t.Errorf("subarraySum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
