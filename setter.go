package roc

import (
	"time"
)

// Setter configures a cache.
type Setter func(c *Cache) error

// WithBucketNum set the number of buckets
func WithBucketNum(num int) Setter {
	return func(c *Cache) error {
		c.BucketNum = num
		return nil
	}
}

// WithGCInterval sets the gc cycle.
func WithGCInterval(interval time.Duration) Setter {
	return func(c *Cache) error {
		c.GCInterval = interval
		return nil
	}
}
