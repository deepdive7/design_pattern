package creational

import (
	"fmt"
	"testing"
	"time"
)

func TestKing(t *testing.T) {
	for i := 0; i < 10; i++ {
		king := GetKing()
		go king.Older()
	}
	time.Sleep(1 * time.Millisecond)
	king := GetKing()
	king.Info()
}

func TestChanKing(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			king := GetKing()
			king.Older()
			seq := <-KingReady
			fmt.Println("seq = ", seq)
		}()
	}
	time.Sleep(1 * time.Millisecond)
	king := GetKing()
	king.Info()
}
