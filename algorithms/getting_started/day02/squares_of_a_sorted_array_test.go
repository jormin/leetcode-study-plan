package day02

import (
	"reflect"
	"testing"
)

// TestSortedSquares 测试有序数组的平方
func TestSortedSquares(t *testing.T) {
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
			args: args{nums: []int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "02",
			args: args{nums: []int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
