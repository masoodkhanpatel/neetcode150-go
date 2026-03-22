package linked_list

import "testing"

func TestLRUCache(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	if val := lru.Get(1); val != 1 {
		t.Errorf("expected 1, got %d", val)
	}
	lru.Put(3, 3) // evicts 2
	if val := lru.Get(2); val != -1 {
		t.Errorf("expected -1, got %d", val)
	}
	lru.Put(4, 4) // evicts 1
	if val := lru.Get(1); val != -1 {
		t.Errorf("expected -1, got %d", val)
	}
	if val := lru.Get(3); val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
	if val := lru.Get(4); val != 4 {
		t.Errorf("expected 4, got %d", val)
	}
}
