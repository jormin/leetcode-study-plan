package day07

import (
	"reflect"
	"testing"
)

// TestPartitionLabels 测试划分字母区间
func TestPartitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{s: "ababcbacadefegdehijhklij"},
			want: []int{9, 7, 8},
		},
		{
			name: "02",
			args: args{s: ""},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
