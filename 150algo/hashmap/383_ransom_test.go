package hashmap

import "testing"

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		desc       string
		ransomNote string
		magazine   string
		result     bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			result:     false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			result:     false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			result:     true,
		},
	}
	for _, tC := range testCases {
		t.Run("tests", func(t *testing.T) {
			if got := canConstruct(tC.ransomNote, tC.magazine); got != tC.result {
				t.Errorf("canConstruct() = %v, want %v", got, tC.result)
			}
		})
	}
}
