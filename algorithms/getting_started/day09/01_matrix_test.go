package day09

import (
	"reflect"
	"testing"
)

// TestUpdateMatrix 测试01 矩阵
func TestUpdateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "02",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
		{
			name: "03",
			args: args{
				mat: [][]int{
					{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
					{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
					{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
					{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
					{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
					{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
					{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
					{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
					{1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
				{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
				{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
				{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
				{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
				{3, 2, 2, 1, 0, 1, 0, 0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
