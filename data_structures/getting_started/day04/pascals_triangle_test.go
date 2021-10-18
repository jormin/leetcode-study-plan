package day04

import (
	"reflect"
	"testing"
)

// 测试杨辉三角
func TestGenerate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{numRows: 5},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name: "02",
			args: args{numRows: 1},
			want: [][]int{
				{1},
			},
		},
		{
			name: "03",
			args: args{numRows: 0},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("generate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
