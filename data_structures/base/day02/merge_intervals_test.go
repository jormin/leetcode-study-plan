package day02

import (
	"reflect"
	"testing"
)

// TestMerge 测试合并区间
func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{
				intervals: [][]int{
					{1, 3}, {2, 6}, {8, 10}, {15, 18},
				},
			},
			want: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
		{
			name: "02",
			args: args{
				intervals: [][]int{
					{1, 4}, {4, 5},
				},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "03",
			args: args{
				intervals: [][]int{
					{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2},
				},
			},
			want: [][]int{
				{0, 9}, {10, 11}, {12, 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("merge() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
