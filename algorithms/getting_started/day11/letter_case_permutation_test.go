package day11

import (
	"reflect"
	"testing"
)

// TestLetterCasePermutation 测试字母大小写全排列
func TestLetterCasePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "01",
			args: args{s: "a1b2"},
			want: []string{"a1b2", "A1b2", "a1B2", "A1B2"},
		},
		{
			name: "02",
			args: args{s: "3z4"},
			want: []string{"3z4", "3Z4"},
		},
		{
			name: "03",
			args: args{s: "12345"},
			want: []string{"12345"},
		},
		{
			name: "04",
			args: args{s: "C"},
			want: []string{"c", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := letterCasePermutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
