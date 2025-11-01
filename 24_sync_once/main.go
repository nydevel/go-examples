package main

import (
	"sync"
	"sync/atomic"
)

type SyncOnce struct {
	done int32
	mut  sync.Mutex
}

func (s *SyncOnce) Do(f func()) {
	if atomic.LoadInt32(&s.done) == 0 {
		s.doSlow(f)
	}
}

func (s *SyncOnce) doSlow(f func()) {
	defer s.mut.Unlock()
	s.mut.Lock()

	if s.done == 0 {
		defer atomic.StoreInt32(&s.done, 1)
		f()
	}
}

func main() {
	var once SyncOnce
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			once.Do(func() {
				println("Hello")
			})
		}()
	}

	wg.Wait()
}
