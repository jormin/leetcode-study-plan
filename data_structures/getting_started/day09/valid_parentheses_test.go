package day09

import "testing"

// TestIsValid 测试有效的括号
func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "02",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "03",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "04",
			args: args{s: "([)]"},
			want: false,
		},
		{
			name: "05",
			args: args{s: "{[]}"},
			want: true,
		},
		{
			name: "06",
			args: args{s: "{"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isValid(tt.args.s); got != tt.want {
					t.Errorf("isValid() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
