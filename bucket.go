package roc

import (
	"container/list"
	"sync"
	"time"
)

type Bucket struct {
	guard sync.Mutex
	objs  *list.List
	coll  map[string]*list.Element
}

func NewBucket() (*Bucket, error) {
	b := new(Bucket)
	return b, nil
}

func (b *Bucket) Get(key string) ([]byte, error) {
	b.guard.Lock()
	defer b.guard.Unlock()
	element, hit := b.coll[key]
	if !hit {
		return nil, ErrMiss
	}
	u, ok := element.Value.(*Unit)
	if !ok {
		return nil, ErrMiss
	}

	if u.Expire() {
		return nil, ErrMiss
	}
	b.objs.MoveToFront(element)
	return u.Data, nil
}

func (b *Bucket) Set(key string, data []byte, d time.Duration) error {
	b.guard.Lock()
	defer b.guard.Unlock()
	element, hit := b.coll[key]
	if !hit {
		unit := new(Unit)
		unit.Key = key
		unit.Data = data
		unit.ExpirationTime = time.Now().UTC().Add(d)
		e := b.objs.PushFront(unit)
		b.coll[key] = e
		return nil
	}
	b.objs.MoveToFront(element)
	unit, ok := element.Value.(*Unit)
	if !ok {
		return nil
	}
	unit.Data = data
	unit.ExpirationTime = time.Now().UTC().Add(d)
	return nil
}

func (b *Bucket) Del(key string) {
	b.guard.Lock()
	defer b.guard.Unlock()
	element, hit := b.coll[key]
	if !hit {
		return
	}
	b.del(key, element)
}

func (b *Bucket) del(key string, e *list.Element) {
	b.objs.Remove(e)
	delete(b.coll, key)
}

func (b *Bucket) gc() {
	b.guard.Lock()
	defer b.guard.Unlock()

	e := b.objs.Back()
	for e != nil {
		unit := e.Value.(*Unit)
		if unit.Expire() {
			b.del(unit.Key, e)
		}
		e = e.Prev()
	}
}
