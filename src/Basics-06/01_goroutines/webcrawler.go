package main

import (
	"fmt"
	"sync"
	"time"
)

const syncParsingSleep, asyncParsingSleep = 2, 2

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	startTime := time.Now()
	Crawl("https://golang.org/", 4, fetcher)
	duration := time.Since(startTime)
	fmt.Printf("\rCrawled in %fs\n", duration.Seconds())
	//found: https://golang.org/ "The Go Programming Language" [https://golang.org/pkg/ https://golang.org/cmd/]
	//found: https://golang.org/pkg/ "Packages" [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
	//found: https://golang.org/ "The Go Programming Language" [https://golang.org/pkg/ https://golang.org/cmd/]
	//found: https://golang.org/pkg/ "Packages" [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
	//not found: https://golang.org/cmd/
	//not found: https://golang.org/cmd/
	//found: https://golang.org/pkg/fmt/ "Package fmt" [https://golang.org/ https://golang.org/pkg/]
	//found: https://golang.org/ "The Go Programming Language" [https://golang.org/pkg/ https://golang.org/cmd/]
	//found: https://golang.org/pkg/ "Packages" [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
	//found: https://golang.org/pkg/os/ "Package os" [https://golang.org/ https://golang.org/pkg/]
	//found: https://golang.org/ "The Go Programming Language" [https://golang.org/pkg/ https://golang.org/cmd/]
	//found: https://golang.org/pkg/ "Packages" [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
	//not found: https://golang.org/cmd/
	// Synchronously
	// Crawled in 26.012577s

	fmt.Println("-----------------------------")

	startTime = time.Now()
	crawler := NewConcurrentCrawler(fetcher, 4, 3)
	crawler.Crawl("https://golang.org/")
	duration = time.Since(startTime)
	fmt.Printf("\rCrawled in %fs\n", duration.Seconds())
	//worker 1: fetching "https://golang.org/" - 1st task creates 2 task: pkg & cmd (worker 2 & 3 get it parallel)
	//	found: https://golang.org/ "The Go Programming Language" [https://golang.org/pkg/ https://golang.org/cmd/]
	//worker 2: fetching "https://golang.org/pkg/" - parallel work of worker 2 & 3
	//	found: https://golang.org/pkg/ "Packages" [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
	// worker 2 has created new jobs `fmt` & `pkg/os` (worker 1 & 2 - after finish this task - get it parallel)
	//worker 3: fetching "https://golang.org/cmd/"
	//	error: not found: https://golang.org/cmd/
	// worker 3 hasn't created new jobs
	//worker 1: fetching "https://golang.org/pkg/fmt/"
	//worker 2: fetching "https://golang.org/pkg/os/"
	//	found: https://golang.org/pkg/fmt/ "Package fmt" [https://golang.org/ https://golang.org/pkg/]
	//	found: https://golang.org/pkg/os/ "Package os" [https://golang.org/ https://golang.org/pkg/]
	// all pages has been crawled

	// Asynchronously
	// Crawled in 6.002301s
}

//////////////////////////////////////////

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// need: Fetch URLs in parallel.
	// need: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	// imitation of connection and processing
	time.Sleep(syncParsingSleep * time.Second)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q %v\n", url, body, urls)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

//////////////////////////////////////////

type crawlJob struct {
	url   string
	depth int
}

type crawlResult struct {
	url   string
	body  string
	urls  []string
	depth int
	err   error
}

// ConcurrentCrawler manages parallel crawling
type ConcurrentCrawler struct {
	fetcher    Fetcher
	maxDepth   int
	maxWorkers int
	visited    *sync.Map
	jobs       chan crawlJob
	results    chan crawlResult
	wg         sync.WaitGroup
}

func NewConcurrentCrawler(fetcher Fetcher, maxDepth, maxWorkers int) *ConcurrentCrawler {
	return &ConcurrentCrawler{
		fetcher:    fetcher,
		maxDepth:   maxDepth,
		maxWorkers: maxWorkers,
		visited:    &sync.Map{},
		jobs:       make(chan crawlJob, 100),
		results:    make(chan crawlResult, 100),
	}
}

// worker handles task from `jobs` channel
func (c *ConcurrentCrawler) worker(id int) {
	for job := range c.jobs {
		// imitation of connection and processing
		time.Sleep(asyncParsingSleep * time.Second)
		fmt.Printf("worker %d: fetching %q\n", id, job.url)

		body, urls, err := c.fetcher.Fetch(job.url)

		c.results <- crawlResult{
			url:   job.url,
			body:  body,
			urls:  urls,
			depth: job.depth,
			err:   err,
		}
	}
}

// processResults handles results and creates new tasks
func (c *ConcurrentCrawler) processResults() {
	for result := range c.results {
		if result.err != nil {
			fmt.Printf("\terror: %v\n", result.err)
		} else {
			fmt.Printf("\tfound: %s %q %v\n", result.url, result.body, result.urls)

			if result.depth > 0 {
				for _, u := range result.urls {
					// check is URL already visited or set true if new one
					if _, loaded := c.visited.LoadOrStore(u, true); loaded {
						continue
					}

					c.wg.Add(1)
					c.jobs <- crawlJob{
						url:   u,
						depth: result.depth - 1,
					}
				}
			}
		}
		c.wg.Done()
	}
}

// Crawl starts parallel crawling
func (c *ConcurrentCrawler) Crawl(startURL string) {
	// start workers in N goroutines
	for i := 1; i <= c.maxWorkers; i++ {
		go c.worker(i)
	}

	// start result handler
	go c.processResults()

	// start with startURL and create first job
	c.wg.Add(1)
	c.visited.LoadOrStore(startURL, true)
	// 1st worker gets this task
	c.jobs <- crawlJob{
		url:   startURL,
		depth: c.maxDepth,
	}

	// waiting for all tasks finishing - then program goes synchronously
	c.wg.Wait()

	close(c.jobs)
	close(c.results)
}
