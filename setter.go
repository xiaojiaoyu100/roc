package roc

import (
	"time"
)

type Setter func(c *Cache) error

func isPowerOfTwo(num int) bool {
	return (num != 0) && ((num & (num - 1)) == 0)
}

func WithBucketNum(num int) Setter {
	return func(c *Cache) error {
		c.BucketNum = num
		return nil
	}
}

func WithGCInterval(interval time.Duration) Setter {
	return func(c *Cache) error {
		c.GCInterval = interval
		return nil
	}
}
