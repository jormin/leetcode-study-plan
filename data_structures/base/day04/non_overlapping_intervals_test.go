package day04

import "testing"

// TestEraseOverlapIntervals 测试无重叠区间
func TestEraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{1, 3},
				},
			},
			want: 1,
		},
		{
			name: "02",
			args: args{
				intervals: [][]int{
					{1, 2},
					{1, 2},
					{1, 2},
				},
			},
			want: 2,
		},
		{
			name: "03",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: 0,
		},
		{
			name: "04",
			args: args{
				intervals: [][]int{
					{1, 100},
					{11, 22},
					{1, 11},
					{2, 12},
				},
			},
			want: 2,
		},
		{
			name: "05",
			args: args{
				intervals: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
					t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
