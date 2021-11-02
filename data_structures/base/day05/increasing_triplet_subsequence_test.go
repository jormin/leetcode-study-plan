package day05

import "testing"

// TestIncreasingTriplet 测试递增的三元子序列
func TestIncreasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "02",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: false,
		},
		{
			name: "03",
			args: args{nums: []int{2, 1, 5, 0, 4, 6}},
			want: true,
		},
		{
			name: "04",
			args: args{nums: []int{2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := increasingTriplet(tt.args.nums); got != tt.want {
					t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
