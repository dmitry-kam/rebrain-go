package main

import (
	cachePackage "04task/cache"
	"fmt"
	"sync"
	"time"
)

type ICache interface {
	Set(key string, value int)
	Increase(key string, value int)
	Get(key string) int
	Remove(key string)
}

const (
	k1   = "key1"
	step = 7
)

func main() {
	var wg sync.WaitGroup
	cache := cachePackage.NewSafeCache()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increaseCacheKey(cache, &wg)
	}

	for i := 0; i < 10; i++ {
		//i := i // copy variable
		wg.Add(1)
		go setCacheKey(cache, i, &wg)
	}

	wg.Wait()
	fmt.Println(cache.Get(k1))
}

func increaseCacheKey(cache ICache, wg *sync.WaitGroup) {
	defer wg.Done()
	cache.Increase(k1, step)
	time.Sleep(time.Millisecond * 100)
}

func setCacheKey(cache ICache, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	cache.Set(k1, step*i)
	time.Sleep(time.Millisecond * 100)
}
