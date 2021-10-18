package day04

import (
	"reflect"
	"testing"
)

// TestMatrixReshape 测试重塑矩阵
func TestMatrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   1,
				c:   4,
			},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "02",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   2,
				c:   4,
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := matrixReshape(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
