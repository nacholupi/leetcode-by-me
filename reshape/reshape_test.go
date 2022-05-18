package reshape

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "zero",
			args: args{
				mat: [][]int{},
				r:   0,
				c:   0,
			},
			want: [][]int{},
		},
		{
			name: "1X1 to 1x1",
			args: args{
				mat: [][]int{{123}},
				r:   1,
				c:   1,
			},
			want: [][]int{{123}},
		},
		{
			name: "2X2 to 1x4",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   1,
				c:   4,
			},
			want: [][]int{{1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
