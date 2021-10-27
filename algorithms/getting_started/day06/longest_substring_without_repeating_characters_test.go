package day06

import "testing"

// TestLengthOfLongestSubstring 测试无重复字符的最长子串
func TestLengthOfLongestSubstring(t *testing.T) {
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
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "02",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "03",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "04",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "05",
			args: args{s: " "},
			want: 1,
		},
		{
			name: "06",
			args: args{s: "dvdf"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
					t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
