package array

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		input    string
		numRows  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
		{"AB", 2, "AB"},
	}

	for _, test := range tests {
		result := convert(test.input, test.numRows)
		if result != test.expected {
			t.Errorf("convert(%q, %d) = %q; expected %q", test.input, test.numRows, result, test.expected)
		}
	}
}
