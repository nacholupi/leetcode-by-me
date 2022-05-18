package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
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
			name: "zero",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "nothing to delete",
			args: args{
				head: &ListNode{
					Val:  2,
					Next: nil},
				val: 0,
			},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			name: "delete everything",
			args: args{
				head: &ListNode{
					Val:  2,
					Next: nil},
				val: 2,
			},
			want: nil,
		},
		{
			name: "delete lastone",
			args: args{
				head: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
				val: 3,
			},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			name: "delete middle node",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				val: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				head: nil,
			},
			want: false,
		}, {
			name: "two nodes",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: false,
		}, {
			name: "two nodes",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: false,
		}, {
			name: "two nodes",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: false,
		}, {
			name: "four nodes",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
