package day13

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{n: 1},
			want: true,
		},
		{
			name: "02",
			args: args{n: 16},
			want: true,
		},
		{
			name: "03",
			args: args{n: 3},
			want: false,
		},
		{
			name: "04",
			args: args{n: 4},
			want: true,
		},
		{
			name: "05",
			args: args{n: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isPowerOfTwo(tt.args.n); got != tt.want {
					t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
