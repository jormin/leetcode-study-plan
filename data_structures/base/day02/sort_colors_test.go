package day02

import (
	"reflect"
	"testing"
)

// TestSortColors 测试颜色分类
func TestSortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "02",
			args: args{nums: []int{2, 0, 1}},
			want: []int{0, 1, 2},
		},
		{
			name: "03",
			args: args{nums: []int{0}},
			want: []int{0},
		},
		{
			name: "04",
			args: args{nums: []int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				sortColors(tt.args.nums)
				if !reflect.DeepEqual(tt.args.nums, tt.want) {
					t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
				}
			},
		)
	}
}
