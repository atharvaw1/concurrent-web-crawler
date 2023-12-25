package worker

import (
	"main/crawler"
	"main/data_store"
	"sync"
	"time"
)

func Work(maxDepth int,urlSet *data_store.URLSet, wg *sync.WaitGroup, ch chan crawler.Page) {
	defer wg.Done()
	noWorkTime := 0
	for {
		select {
		case page   := <-ch:
			noWorkTime = 0
			wg.Add(1)
			go crawler.Crawl(page,maxDepth,urlSet, wg, ch)
		default:
			time.Sleep(100 * time.Millisecond)
			//IF no work done in 10 seconds then exit
			noWorkTime++
			if noWorkTime > 50 {
				return
			}
		}


	}	
}

