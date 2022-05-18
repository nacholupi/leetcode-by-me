package pascal

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one case",
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "two rows case",
			args: args{
				numRows: 2,
			},
			want: [][]int{{1}, {1, 1}},
		},
		{
			name: "three rows case",
			args: args{
				numRows: 3,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name: "four rows case",
			args: args{
				numRows: 4,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			name: "five rows case",
			args: args{
				numRows: 5,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
