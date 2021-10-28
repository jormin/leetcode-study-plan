package day07

import (
	"reflect"
	"testing"
)

// TestFloodFill 测试图像渲染
func TestFloodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "02",
			args: args{
				image:    nil,
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: nil,
		},
		{
			name: "03",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 1, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 1,
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("floodFill() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
