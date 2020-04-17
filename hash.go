package roc

import "hash/fnv"

func (c *Cache) hashIndex(key string) (int, error) {
	hash := fnv.New64a()
	_, err := hash.Write([]byte(key))
	if err != nil {
		return 0, err
	}
	return int(hash.Sum64() & uint64(c.BucketNum-1)), nil
}
