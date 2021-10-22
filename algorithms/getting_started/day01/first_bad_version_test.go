package day01

import "testing"

// TestFirstBadVersion 测试第一个错误的版本
func TestFirstBadVersion(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{n: 5, bad: 4},
			want: 4,
		},
		{
			name: "02",
			args: args{n: 1, bad: 1},
			want: 1,
		},
		{
			name: "03",
			args: args{n: 5, bad: 2},
			want: 2,
		},
		{
			name: "04",
			args: args{n: 6, bad: 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				setBadVersion(tt.args.bad)
				if got := firstBadVersion(tt.args.n); got != tt.want {
					t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// TestIsBadVersion 测试是否错误版本
func TestIsBadVersion(t *testing.T) {
	type args struct {
		bad     int
		version int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				bad:     1,
				version: 1,
			},
			want: true,
		},
		{
			name: "02",
			args: args{
				bad:     2,
				version: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				setBadVersion(tt.args.bad)
				if got := isBadVersion(tt.args.version); got != tt.want {
					t.Errorf("isBadVersion() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// TestSetBadVersion 测试设置错误版本
func TestSetBadVersion(t *testing.T) {
	type args struct {
		version int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "01",
			args: args{version: 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				setBadVersion(tt.args.version)
			},
		)
	}
}
