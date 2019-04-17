package cache

import "testing"

func TestCache_GetSize(t *testing.T) {

	// given
	c := InitializeCache()

	// expect
	if c.GetSize() != 0 {
		t.Error("Cache is not empty after initialization")
	}

	// when
	c.PutValue("one", "one")
	c.PutValue("two", "two")

	// expect
	if c.GetSize() != 2 {
		t.Errorf("Cache size is expected to be %d, got %d", 2, c.GetSize())
	}

	// when
	c.DeleteValue("one")

	// expect
	if c.GetSize() != 1 {
		t.Errorf("Cache size is expected to be %d, got %d", 1, c.GetSize())
	}
}
