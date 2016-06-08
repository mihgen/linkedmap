package linkedmap

import "testing"

func TestElementKeyUpdated(t *testing.T) {
	lm := New()
	key := "key"
	lm.Add(key, "value")

	if lm.last.key != key {
		t.Errorf("Expected element's key %v, got %v", key, lm.last.key)
	}
}
