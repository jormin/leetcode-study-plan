package day06

import "testing"

// TestAddStrings 测试字符串相加
func TestAddStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		},
		{
			name: "02",
			args: args{
				num1: "456",
				num2: "77",
			},
			want: "533",
		},
		{
			name: "03",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "04",
			args: args{
				num1: "9",
				num2: "9",
			},
			want: "18",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
					t.Errorf("addStrings() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
