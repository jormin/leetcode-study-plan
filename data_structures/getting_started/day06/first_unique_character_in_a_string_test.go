package day06

import "testing"

// TestFirstUniqChar 测试字符串中的第一个唯一字符
func TestFirstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{s: "leetcode"},
			want: 0,
		},
		{
			name: "02",
			args: args{s: "loveleetcode"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := firstUniqChar(tt.args.s); got != tt.want {
					t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
