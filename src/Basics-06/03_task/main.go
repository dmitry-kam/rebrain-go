package main

import (
	cachePackage "03task/cache"
	"fmt"
	"sync"
)

type ICache interface {
	Set(key string, value int)
	Increase(key string, value int)
	Get(key string) int
	Remove(key string)
}

func main() {
	pages := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// `Your pages have been visited 9 times`
	visitors := []int{1}
	// Waiting for 9x9=81 views total, but `fatal error: concurrent map writes`
	//visitors = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup

	fmt.Println("Unsafe cache check")
	unsafeCacheViewsCounter := cachePackage.New()

	for _, visitor := range visitors {
		wg.Add(1)
		go viewAllPages(unsafeCacheViewsCounter, visitor, pages, &wg)
	}

	wg.Wait()

	totalViews := 0
	for _, page := range pages {
		totalViews += unsafeCacheViewsCounter.Get(page)
	}

	fmt.Printf("Your pages have been visited %d times\n", totalViews)

	fmt.Printf("\nThread-safe cache check\n")
	// `Your pages have been visited 81 times`
	visitors = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	safeCacheViewsCounter := cachePackage.NewSafeCache()

	for _, visitor := range visitors {
		wg.Add(1)
		go viewAllPages(safeCacheViewsCounter, visitor, pages, &wg)
	}

	wg.Wait()

	totalViews = 0
	for _, page := range pages {
		totalViews += safeCacheViewsCounter.Get(page)
	}

	fmt.Printf("Your pages have been visited %d times\n", totalViews)
}

func viewAllPages(c ICache, v int, pages []string, wg *sync.WaitGroup) {
	for _, page := range pages {
		c.Increase(page, 1)
	}
	wg.Done()
	fmt.Printf("Visitor %d has visited all the pages\n", v)
}
