package cache

import (
	"testing"
)

func TestSetAndGet(t *testing.T) {
	c := New()

	c.Set("some_key", "some_value")

	value, ok := c.Get("some_key")
	if !ok {
		t.Error("Expected key 'some_key' to be found, but it wasn't")
	}
	if value != "some_value" {
		t.Errorf("Expected value 'some_value', but got %v", value)
	}
}

func TestDelete(t *testing.T) {
	c := New()

	c.Set("some_key", "some_value")

	c.Delete("some_key")
	_, ok := c.Get("some_key")

	if ok {
		t.Error("Expected key 'some_key' to be deleted, but it was found")
	}
}

func TestClear(t *testing.T) {
	c := New()

	c.Set("some_key", "some_value")
	c.Set("some_key_2", "some_value_2")

	c.Clear()

	if _, found := c.Get("some_key"); found {
		t.Error("Expected cache to be cleared, but 'some_key' was found")
	}
	if _, found := c.Get("some_key_2"); found {
		t.Error("Expected cache to be cleared, but 'some_key_2' was found")
	}
}
