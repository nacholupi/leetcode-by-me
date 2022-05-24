package bsearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -1,
		},
		{
			name: "one",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "array len 2 with target",
			args: args{
				nums:   []int{1, 2},
				target: 1,
			},
			want: 0,
		},
		{
			name: "array len 3 with target",
			args: args{
				nums:   []int{1, 2, 3},
				target: 3,
			},
			want: 2,
		},
		{
			name: "array len 3 not containig target",
			args: args{
				nums:   []int{1, 2, 4},
				target: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
