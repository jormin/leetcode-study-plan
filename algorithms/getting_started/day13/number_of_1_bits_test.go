package day13

import "testing"

// TestHammingWeight 测试位1的个数
func TestHammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{num: 00000000000000000000000000001011},
			want: 3,
		},
		{
			name: "02",
			args: args{num: 00000000000000000000000010000000},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := hammingWeight(tt.args.num); got != tt.want {
					t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
