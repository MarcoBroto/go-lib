package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

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

var mutex sync.Mutex

func crawl(url string, depth int, fetcher Fetcher, seen map[string]bool, ch1 chan string) {
	defer close(ch1)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch1 <- fmt.Sprint(err)
		return
	}

	mutex.Lock()
	if ok := seen[url]; !ok {
		seen[url] = true
		ch1 <- fmt.Sprintf("found: %s %q", url, body)
	}
	mutex.Unlock()

	for _, u := range urls {
		ch2 := make(chan string)
		go crawl(u, depth-1, fetcher, seen, ch2)
		for v := range ch2 {
			ch1 <- v
		}
	}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan string)
	go crawl(url, depth, fetcher, make(map[string]bool), ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func TestWebCrawler() { // main
	Crawl("https://golang.org/", 4, fetcher)
}
