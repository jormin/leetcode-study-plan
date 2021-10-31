package day03

import (
	"reflect"
	"testing"
)

// TestGenerateMatrix 测试螺旋矩阵 II
func TestGenerateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{n: 3},
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "02",
			args: args{n: 1},
			want: [][]int{
				{1},
			},
		},
		{
			name: "03",
			args: args{n: 4},
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
