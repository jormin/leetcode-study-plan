package day06

import "testing"

// TestIsAnagram 测试有效的字母异位词
func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "02",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
					t.Errorf("isAnagram() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
