package day05

import (
	"reflect"
	"testing"
)

// TestProductExceptSelf 测试除自身以外数组的乘积
func TestProductExceptSelf(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
