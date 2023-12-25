package main

import (
	"main/crawler"
	"main/worker"
	"main/data_store"
	"sync"
	"fmt"
)

const (
	maxWorkers = 5
	maxDepth = 2
)

func main() {
	var wg sync.WaitGroup

	urlSet := data_store.NewURLSet()
	startURL := "http://en.wikipedia.org/wiki/Main_Page"

	urlCh := make(chan crawler.Page , 10)

	// Crawl the start URL
	wg.Add(1)
	go crawler.Crawl(crawler.Page{startURL, 1},maxDepth, urlSet,  &wg, urlCh)

	fmt.Println("Crawling...")
	// Start workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker.Work(maxDepth,urlSet, &wg, urlCh)
	}
	fmt.Println("Workers started...")
	wg.Wait()
	fmt.Println("Crawling complete")
	close(urlCh)
}
