package day01

import "testing"

// TestSearchInsert 测试搜索插入位置
func TestSearchInsert(t *testing.T) {
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
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "02",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "03",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "04",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "05",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "06",
			args: args{
				nums:   []int{1,3},
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("searchInsert() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
