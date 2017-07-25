package main

import (
	"fmt"
	"sync"
	"time"
)

type safeVariable struct {
	sync.Mutex
	count int
}

func (s *safeVariable) Plus() {
	s.Lock()
	defer s.Unlock()

	s.count++
}

func (s *safeVariable) Count() int {
	return s.count
}

func main() {
	s := &safeVariable{}

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			s.Plus()
			wg.Done()
		}()
	}

	wg.Wait()

	// ba wait...
	fmt.Println("with wait", s.Count())

	s.count = 0

	for i := 0; i < 10; i++ {
		go s.Plus()
	}

	// bedone wait vali sabr mikonim ta tamame go ha kareshon tamom she
	time.Sleep(time.Second)
	fmt.Println("with second", s.Count())
}
