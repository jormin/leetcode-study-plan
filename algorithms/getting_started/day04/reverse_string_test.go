package day04

import (
	"reflect"
	"testing"
)

// TestReverseString 测试反转字符串
func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "01",
			args: args{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "02",
			args: args{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				reverseString(tt.args.s)
				if !reflect.DeepEqual(tt.args.s, tt.want) {
					t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
				}
			},
		)
	}
}
