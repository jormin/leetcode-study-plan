package day01

import "testing"

// TestSearch 测试二分查找
func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				nums:   nil,
				target: 0,
			},
			want: -1,
		},
		{
			name: "02",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "03",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "04",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "05",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 13,
			},
			want: -1,
		},
		{
			name: "06",
			args: args{
				nums:   []int{2, 5},
				target: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := search(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("search() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
