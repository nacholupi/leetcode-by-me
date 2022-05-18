package reverselist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "zero case",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "one case",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "len two case",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
		{
			name: "len three case",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
