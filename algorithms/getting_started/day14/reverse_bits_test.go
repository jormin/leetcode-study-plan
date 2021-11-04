package day14

import "testing"

// 测试颠倒二进制位
func TestReverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "01",
			args: args{num: 43261596},
			want: 964176192,
		},
		{
			name: "02",
			args: args{num: 4294967293},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := reverseBits(tt.args.num); got != tt.want {
					t.Errorf("reverseBits() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
