package back

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits   string
		expected []string
	}{
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
		{"23", []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"}},
		{"7", []string{"p", "q", "r", "s"}},
		{"79", []string{"pw", "qw", "rw", "sw", "px", "qx", "rx", "sx", "py", "qy", "ry", "sy", "pz", "qz", "rz", "sz"}},
	}

	for _, test := range tests {
		result := letterCombinations(test.digits)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For digits %q, expected %v, but got %v", test.digits, test.expected, result)
		}
	}
}
