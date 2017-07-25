package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var u = uint32(0)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddUint32(&u, 1)
			time.Sleep(time.Millisecond)
		}()
	}
	time.Sleep(time.Second)

	counter := atomic.LoadUint32(&u)
	fmt.Println("counter:", counter)
}
