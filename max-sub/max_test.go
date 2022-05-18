package max

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "one",
			args: args{nums: []int{2}},
			want: 2,
		},
		{
			name: "many",
			args: args{nums: []int{2, 1, 3}},
			want: 6,
		},
		{
			name: "leetcode 1st case",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "variation of leetcode 1st case",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 20}},
			want: 21,
		},
		{
			name: "variation two of leetcode 1st case",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -7, 20}},
			want: 20,
		},
		{
			name: "leetcode 2nd case",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
