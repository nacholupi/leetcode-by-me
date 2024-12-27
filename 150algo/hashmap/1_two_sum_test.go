package hashmap

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	type _case struct {
		name string
		args args
		want []int
	}
	tests := []_case{
		{
			name: "first two index",
			args: args{[]int{1, 2}, 3},
			want: []int{0, 1},
		},
		{
			name: "nums len four",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{0, 1},
		},
		{
			name: "nums len three",
			args: args{[]int{3, 2, 4}, 6},
			want: []int{1, 2},
		},
		{
			name: "nums len two",
			args: args{[]int{3, 3}, 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
