package main

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/text/language"
)

type request struct {
	from language.Tag
	to   language.Tag
	data string
}

type cachedRequest struct {
	request
	result            string
	expireAtTimestamp int64
}

type localCache struct {
	stop chan struct{}

	wg       sync.WaitGroup
	mu       sync.RWMutex
	requests []cachedRequest
}

func newLocalCache(cleanupInterval time.Duration) *localCache {

	var cachedReqs []cachedRequest
	lc := &localCache{
		requests: cachedReqs,
		stop:     make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		//	lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return lc
}

/* func (lc *localCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for index, cu := range lc.requests {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					remove(lc.requests, index)
				}
			}
			lc.mu.Unlock()
		}
	}
}
*/
func (lc *localCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *localCache) update(u request, result string, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.requests = append(lc.requests, cachedRequest{
		request:           u,
		expireAtTimestamp: expireAtTimestamp,
		result:            result,
	})
}

var (
	errNotInCache = errors.New("the request isn't in cache")
)

func (lc *localCache) read(request request) (cachedRequest, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	for _, element := range lc.requests {
		if element.request == request {
			return element, nil
		}
	}

	return cachedRequest{}, errors.New(errNotInCache.Error())
}

/* func (lc *localCache) delete(id int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.requests, id)
}
*/
