package array

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"throne", "throne"}, "throne"},
		{[]string{"throne", "dungeon"}, ""},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %s but got %s", test.input, test.expected, result)
		}
	}
}
