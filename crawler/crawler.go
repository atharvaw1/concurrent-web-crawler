package crawler

import (
	"net/http"
	"sync"
	"golang.org/x/net/html"
	"main/data_store"
	//"fmt"
) 

type Page struct {
	URL string
	Depth int
}

func Crawl(page Page, maxDepth int, urlSet *data_store.URLSet, wg *sync.WaitGroup,  ch chan Page) {

	defer wg.Done()
	if page.Depth > maxDepth {
		return
	} 
	// Crawl the URL
	resp, err := http.Get(page.URL)
	if err != nil {
		//fmt.Println("Error fetching URL: ", err)
		return
	}
	defer resp.Body.Close()

	// Check content type is HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		//fmt.Println("Error parsing HTML: ", err)
		return
	}

	// Find links
	links := ExtractLinks(doc)
	
	for _, link := range links {
		if urlSet.Add(link) {
			//fmt.Println("Found link: ", link)
			ch <- Page{link, page.Depth + 1}
		}
	}

}
