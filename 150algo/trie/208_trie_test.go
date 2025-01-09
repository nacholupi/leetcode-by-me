package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()

	// Test Insert and Search
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Expected to find 'apple' in the trie")
	}
	if trie.Search("app") {
		t.Errorf("Did not expect to find 'app' in the trie")
	}

	// Test StartsWith
	if !trie.StartsWith("app") {
		t.Errorf("Expected to find prefix 'app' in the trie")
	}

	// Test Insert and Search for a new word
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Expected to find 'app' in the trie after insertion")
	}

	// Test Search for a word that does not exist
	if trie.Search("banana") {
		t.Errorf("Did not expect to find 'banana' in the trie")
	}

	// Test StartsWith for a prefix that does not exist
	if trie.StartsWith("ban") {
		t.Errorf("Did not expect to find prefix 'ban' in the trie")
	}
}
