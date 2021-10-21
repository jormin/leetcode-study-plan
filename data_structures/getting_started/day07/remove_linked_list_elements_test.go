package day07

import (
	"reflect"
	"testing"
)

// TestRemoveElements 测试移除链表元素
func TestRemoveElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "02",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "03",
			args: args{
				head: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
				val: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("removeElements() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
