package hashmap

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"abc", "dog cat fish", true},
		{"", "", true},
		{"a", "dog", true},
		{"a", "dog cat", false},
	}

	for _, tt := range tests {
		t.Run(tt.pattern+"_"+tt.s, func(t *testing.T) {
			if got := wordPattern(tt.pattern, tt.s); got != tt.want {
				t.Errorf("wordPattern(%q, %q) = %v, want %v", tt.pattern, tt.s, got, tt.want)
			}
		})
	}
}
