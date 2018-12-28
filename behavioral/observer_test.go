package behavioral

import (
	"sync"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	sub := &Subject{
		make(map[int]chan<- int, 100),
		0,
		sync.Mutex{},
	}
	N := 5
	observers := make([]Observer, N, N)
	for i := 0; i < N; i++ {
		a := i
		observers[i] = &BinaryObserver{
			index:      a,
			localState: 0,
		}
		observers[i].subscribe(sub)
		go func(k int) {
			observers[k].handle()
		}(i)
	}
	sub.UpdateState(20)
	time.Sleep(1 * time.Second)
}
