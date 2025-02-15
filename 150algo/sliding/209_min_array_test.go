package sliding

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{15, []int{1, 2, 3, 4, 5}, 5},
		{100, []int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d; want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
