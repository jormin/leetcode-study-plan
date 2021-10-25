package day04

import "testing"

// TestReverseWords 测试反转字符串中的单词 III
func TestReverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{s: ""},
			want: "",
		},
		{
			name: "02",
			args: args{s: "Let's take LeetCode contest"},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := reverseWords(tt.args.s); got != tt.want {
					t.Errorf("reverseWords() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
