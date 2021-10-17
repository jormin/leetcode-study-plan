package day03

import "testing"

// TestMaxProfit 测试买卖股票的最佳时机
func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "02",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "03",
			args: args{prices: []int{1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := maxProfit(tt.args.prices); got != tt.want {
					t.Errorf("maxProfit() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
