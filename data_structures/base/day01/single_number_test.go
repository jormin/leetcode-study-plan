package day01

import "testing"

// TestSingleNumber 测试只出现一次的数字
func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{nums: []int{2, 2, 1}},
			want: 1,
		},
		{
			name: "02",
			args: args{nums: []int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "03",
			args: args{nums: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := singleNumber(tt.args.nums); got != tt.want {
					t.Errorf("singleNumber() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
