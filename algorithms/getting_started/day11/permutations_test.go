package day11

import (
	"reflect"
	"testing"
)

// TestPermute 测试全排列
func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "02",
			args: args{nums: []int{0, 1}},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "03",
			args: args{nums: []int{1}},
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("permute() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
