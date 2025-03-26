package linked

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	// Test Put and Get
	cache.Put(1, 1)
	cache.Put(2, 2)
	if val := cache.Get(1); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Test capacity limit and eviction
	cache.Put(3, 3) // Evicts key 2
	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}

	// Test adding new key and eviction
	cache.Put(4, 4) // Evicts key 1
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}
	if val := cache.Get(3); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if val := cache.Get(4); val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}

	// Test updating an existing key
	cache.Put(4, 40)
	if val := cache.Get(4); val != 40 {
		t.Errorf("Expected 40, got %d", val)
	}
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	if val := cache.Get(2); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	cache.Put(4, 1)
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}
	if val := cache.Get(2); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
}
