package kdiff

import (
	"testing"
)

func TestCountKDifference(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	type _case struct {
		name string
		args args
		want int
	}
	tests := []_case{
		{
			name: "nums.lengh 1",
			args: args{[]int{3}, 6},
			want: 0,
		},
		{
			name: "0 pairs with diff",
			args: args{[]int{3, 2}, 2},
			want: 0,
		},
		{
			name: "1 pair with diff",
			args: args{[]int{1, 2}, 1},
			want: 1,
		},
		{
			name: "4 pairs with diff",
			args: args{[]int{2, 1, 1, 2}, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountKDifference(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("CountKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
