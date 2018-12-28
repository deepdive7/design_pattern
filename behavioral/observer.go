package behavioral

import (
	"fmt"
	"sync"
	"time"
)

type Subject struct {
	observerChans map[int]chan<- int
	state         int
	mu            sync.Mutex
}

func (s *Subject) AddObserver(index int, observer chan<- int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observerChans[index] = observer
}

func (s *Subject) NotifyAll() {
	for i, c := range s.observerChans {
		go func(k int, ch chan<- int) {
			fmt.Println("Sent to ", k)
			ch <- s.state
		}(i, c)
	}
	time.Sleep(100 * time.Millisecond)
}

func (s *Subject) UpdateState(newState int) {
	s.state = newState
	s.NotifyAll()
}

type Observer interface {
	update(int)
	subscribe(*Subject)
	handle()
}

type BinaryObserver struct {
	index      int
	localState int
	r          <-chan int
}

func (b *BinaryObserver) update(newState int) {
	b.localState = newState
	fmt.Printf("Index %d, New State: %d\n", b.index, b.localState)
}

func (b *BinaryObserver) subscribe(sub *Subject) {
	newC := make(chan int)
	b.r = (<-chan int)(newC)
	sub.AddObserver(b.index, newC)
}

func (b *BinaryObserver) handle() {
	var newState int
	for {
		newState = <-b.r
		go b.update(newState)
	}
}
