package day01

import "testing"

// TestContainsDuplicate 测试存在重复元素
func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{nums: []int{1,2,3,1}},
			want: true,
		},
		{
			name: "02",
			args: args{nums: []int{1,2,3,4}},
			want: false,
		},
		{
			name: "03",
			args: args{nums: []int{1,1,1,3,3,4,3,2,4,2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := containsDuplicate(tt.args.nums); got != tt.want {
					t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
