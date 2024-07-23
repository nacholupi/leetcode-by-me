package array

import (
	"reflect"
	"testing"
)

func Test_removeDup(t *testing.T) {

	type args struct {
		nums []int
	}

	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "case O: one",
			args: args{
				nums: []int{2},
			},
			want:     1,
			wantNums: []int{2},
		},
		{
			name: "case M many",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:     2,
			wantNums: []int{1, 2, 2},
		},
		{
			name: "case M many 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.nums, tt.wantNums) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums, tt.wantNums)
			}
		})
	}
}
