package day06

import "testing"

// TestCheckInclusion 测试字符串的排列
func TestCheckInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "02",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
		{
			name: "03",
			args: args{
				s1: "ab",
				s2: "bad",
			},
			want: true,
		},
		{
			name: "04",
			args: args{
				s1: "abcd",
				s2: "bad",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
					t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
