package day06

import "testing"

// TestCanConstruct 测试赎金信
func TestCanConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "02",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "03",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
					t.Errorf("canConstruct() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
