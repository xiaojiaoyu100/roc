package roc

import "testing"

func TestBucket_hashIndex(t *testing.T) {
	cache, err := New(WithBucketNum(16))
	if err != nil {
		t.Fatalf("cache err = %+v", err)
	}
	tests := [...]struct {
		In  string
		Out int
	}{
		{
			"",
			5,
		},
		{
			"1323232",
			9,
		},
		{
			".",
			1,
		},
		{
			"nfjndjfndsjfnjdnjfndjscds",
			15,
		},
	}
	for _, tt := range tests {
		idx, err := cache.hashIndex(tt.In)
		if err != nil {
			t.Fatalf("key = %s, err = %#v", tt.In, err)
		}
		if idx != tt.Out {
			t.Fatalf("key = %s, want = %d, got = %d", tt.In, tt.Out, idx)
		}
	}

}
