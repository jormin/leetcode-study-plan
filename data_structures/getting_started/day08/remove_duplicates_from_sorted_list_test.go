package day08

import (
	"reflect"
	"testing"
)

// TestDeleteDuplicates 测试删除排序链表中的重复元素
func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "02",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "03",
			args: args{head: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			}},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "04",
			args: args{head: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			}},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
		{
			name: "05",
			args: args{head: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			}},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
