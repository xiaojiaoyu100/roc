package roc

import (
	"testing"
	"time"
)

func TestCache_Set(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("new cache failed")
	}
	err = c.Set("one", "two", time.Second)
	if err != nil {
		t.Fatalf("set failed")
	}
	t.Log(c.Get("one"))
}
