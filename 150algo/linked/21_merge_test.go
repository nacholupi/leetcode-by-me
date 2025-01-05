package linked

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		first *ListNode
		sec   *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "zero",
			args: args{
				first: NewListNode([]int{}),
				sec:   NewListNode([]int{}),
			},
			want: nil,
		},
		{
			name: "first listNode nil",
			args: args{
				sec: NewListNode([]int{19, 2, 4}),
			},
			want: NewListNode([]int{19, 2, 4}),
		},
		{
			name: "sec listNode nil",
			args: args{
				sec: NewListNode([]int{19, 2, 4}),
			},
			want: NewListNode([]int{19, 2, 4}),
		},
		{
			name: "merge two simple nodes",
			args: args{
				first: NewListNode([]int{19}),
				sec:   NewListNode([]int{19}),
			},
			want: NewListNode([]int{19, 19}),
		},
		{
			name: "merge two list nodes",
			args: args{
				first: NewListNode([]int{1, 2, 4}),
				sec:   NewListNode([]int{1, 3, 4}),
			},
			want: NewListNode([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.first, tt.args.sec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewListNode(arr []int) *ListNode {
	var first *ListNode
	var last *ListNode

	for i, v := range arr {
		nn := &ListNode{Val: v}
		if i == 0 {
			first = nn
		} else {
			last.Next = nn
		}
		last = nn
	}
	return first
}
