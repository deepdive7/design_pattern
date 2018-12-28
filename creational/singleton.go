package creational

import (
	"fmt"
	"sync"
)

type king struct {
	age int
	mu  sync.RWMutex
}

func (k *king) Older() {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.age++
}

// Info ...
func (k *king) Info() {
	fmt.Printf("The king is %d years old!\n", k.age)
}

//GetKing return the only one king
func GetKing() *king {
	if singleKing == nil {
		fmt.Println("Init king once")
		singleKing = &king{0, sync.RWMutex{}}
	}
	return singleKing
}

func ChanKing() *king {
	KingReady <- kingReqSeq
	kingReqSeq++
	if singleKing == nil {
		fmt.Println("Init king once")
		singleKing = &king{0, sync.RWMutex{}}
	}
	return singleKing
}
