package linked

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
