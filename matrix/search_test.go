package matrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "zero",
			args: args{
				matrix: [][]int{},
				target: 0,
			},
			want: false,
		},
		{
			name: "find first",
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
				target: 2,
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				matrix: [][]int{{1, 20}, {40, 50}},
				target: 51,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
