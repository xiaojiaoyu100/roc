# Roc

Roc is a key-value memory cache.

## Feature

* Volatile LRU
* Quick GC

## Usage
```go
package main

import (
	"fmt"
	"github.com/xiaojiaoyu100/roc"
	"time"
)

func main() {
	cache, err := roc.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cache.Set("myfirstkey", "123", time.Second*3); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cache.Get("myfirstkey"))
	fmt.Println(cache.Del("myfirstkey"))
	fmt.Println(cache.Get("myfirstkey"))
}
```


