package roc

import (
	"sync"
	"time"
)

// Bucket is a piece of cache.
type Bucket struct {
	coll sync.Map
}

// NewBucket returns a bucket.
func NewBucket() (*Bucket, error) {
	b := new(Bucket)
	return b, nil
}

// Get returns a corresponding value.
func (b *Bucket) Get(key string) (interface{}, error) {
	value, hit := b.coll.Load(key)
	if !hit {
		return nil, ErrMiss
	}
	u, ok := value.(*Unit)
	if !ok {
		return nil, ErrMiss
	}
	if u.Expire() {
		return nil, ErrMiss
	}
	return u.Data, nil
}

// Set sets a value for a key.
func (b *Bucket) Set(key string, data interface{}, d time.Duration) error {
	value, hit := b.coll.Load(key)
	if !hit {
		unit := new(Unit)
		unit.Key = key
		unit.Data = data
		unit.ExpirationTime = time.Now().UTC().Add(d)
		b.coll.Store(key, unit)
		return nil
	}
	unit, ok := value.(*Unit)
	if !ok {
		return nil
	}
	unit.Data = data
	unit.ExpirationTime = time.Now().UTC().Add(d)
	return nil
}

// Del deletes a key.
func (b *Bucket) Del(key string) {
	b.coll.Delete(key)
}

func (b *Bucket) gc() {
	b.coll.Range(func(key, value interface{}) (keep bool) {
		keep = true
		unit, ok := value.(*Unit)
		if !ok {
			return
		}
		if !unit.Expire() {
			return
		}
		b.Del(key.(string))
		return
	})
}
