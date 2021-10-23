package day03

import (
	"reflect"
	"testing"
)

// TestMoveZeroes 测试移动零
func TestMoveZeroes(t *testing.T) {
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
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				moveZeroes(tt.args.nums)
				if !reflect.DeepEqual(tt.args.nums, tt.want) {
					t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
				}
			},
		)
	}
}
