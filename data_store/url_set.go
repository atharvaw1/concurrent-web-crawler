package data_store

import (
	"sync"
)

type URLSet struct {
	visited map[string]struct{}
	mu sync.Mutex
}
 func NewURLSet() *URLSet {
	return &URLSet{
		visited: make(map[string]struct{}),
	}
}

func (u *URLSet) Add(url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.visited[url]; ok {
		return false
	}
	u.visited[url] = struct{}{}
	return true
}	


