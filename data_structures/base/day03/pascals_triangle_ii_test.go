package day03

import (
	"reflect"
	"testing"
)

// TestGetRow 测试getRow
func TestGetRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{rowIndex: 3},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "02",
			args: args{rowIndex: 0},
			want: []int{1},
		},
		{
			name: "03",
			args: args{rowIndex: 1},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getRow() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
