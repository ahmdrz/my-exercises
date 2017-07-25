package main

import "sync/atomic"
import "fmt"

type once struct {
	done uint64
}

func (o *once) Do(f func()) {
	if atomic.LoadUint64(&o.done) == 1 {
		return
	}

	f()
	atomic.StoreUint64(&o.done, 1)
}

func main() {
	_once := once{}
	for i := 0; i < 10; i++ {
		_once.Do(func() {
			fmt.Println("print hello")
		})
	}
}
