package merges

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums2 empty",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "nums1 empty",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
		{
			name: "merge leetcode example",
			args: args{
				nums1: []int{1, 2, 6, 7, 0, 0, 0},
				m:     4,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 5, 6, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
