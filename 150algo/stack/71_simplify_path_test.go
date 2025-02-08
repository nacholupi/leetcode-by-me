package stack

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "root path",
			input:    "/",
			expected: "/",
		},
		{
			name:     "simple path",
			input:    "/home/",
			expected: "/home",
		},
		{
			name:     "path with double slashes",
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			name:     "path with current directory",
			input:    "/home/./foo/",
			expected: "/home/foo",
		},
		{
			name:     "path with parent directory",
			input:    "/home/../foo/",
			expected: "/foo",
		},
		{
			name:     "complex path",
			input:    "/a/./b/../../c/",
			expected: "/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifyPath(tt.input)
			if result != tt.expected {
				t.Errorf("simplifyPath(%s) = %s; want %s",
					tt.input, result, tt.expected)
			}
		})
	}
}
