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

func BenchmarkBucket_Get(b *testing.B) {
	b.ResetTimer()
	c, err := New()
	if err != nil {
		b.Fatalf("new cache failed")
	}
	for i := 0; i < b.N; i++ {
		c.Get("sssssssssssssssssssssssss")
	}
}
