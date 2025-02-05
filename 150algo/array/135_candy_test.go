package array

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 3, 4, 5, 1}, 11},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 4, 3, 2, 1}, 15},
		{[]int{1, 3, 2, 2, 1}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 45},
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := candy(tt.ratings); got != tt.want {
				t.Errorf("candy(%v) = %v, want %v", tt.ratings, got, tt.want)
			}
		})
	}
}
