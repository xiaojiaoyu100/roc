package roc

import (
	"container/list"
	"context"
	"time"

	"github.com/xiaojiaoyu100/curlew"
)

type Cache struct {
	GCInterval time.Duration
	BucketNum  int
	buckets    []*Bucket
	dispatcher *curlew.Dispatcher
	close      chan struct{}
}

func New(setters ...Setter) (*Cache, error) {
	c := new(Cache)
	c.GCInterval = 60 * time.Second
	c.BucketNum = 16

	for _, setter := range setters {
		if err := setter(c); err != nil {
			return nil, err
		}
	}

	if !isPowerOfTwo(c.BucketNum) {
		return nil, ErrorBucketNum
	}
	c.buckets = make([]*Bucket, 0, c.BucketNum)
	for idx := 0; idx < c.BucketNum; idx++ {
		bucket, err := NewBucket()
		if err != nil {
			return nil, err
		}
		bucket.coll = make(map[string]*list.Element)
		bucket.objs = list.New()
		c.buckets = append(c.buckets, bucket)
	}

	monitor := func(err error) {}

	var err error
	c.dispatcher, err = curlew.New(curlew.WithMonitor(monitor))
	if err != nil {
		return nil, err
	}

	c.close = make(chan struct{})
	c.gc()
	return c, nil
}

func (c *Cache) gc() {
	go func() {
		ticker := time.NewTicker(c.GCInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				for _, bucket := range c.buckets {
					j := curlew.NewJob()
					j.Fn = func(_ context.Context, arg interface{}) error {
						b := arg.(*Bucket)
						b.gc()
						return nil
					}
					j.Arg = bucket
					c.dispatcher.SubmitAsync(j)
				}
			case <-c.close:
				return
			}
		}
	}()
}

func (c *Cache) Get(key string) (interface{}, error) {
	idx, err := c.hashIndex(key)
	if err != nil {
		return nil, err
	}
	return c.buckets[idx].Get(key)
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	idx, err := c.hashIndex(key)
	if err != nil {
		return err
	}
	return c.buckets[idx].Set(key, value, duration)
}

func (c *Cache) Del(key string) error {
	idx, err := c.hashIndex(key)
	if err != nil {
		return err
	}
	c.buckets[idx].Del(key)
	return nil
}
