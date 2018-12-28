package behavioral

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	com := Computer{Num1: 10, Num2: 2}
	strate := NewStrategy("a")
	com.SetStrategy(strate)
	fmt.Println(com.Do())
	strate = NewStrategy("s")
	com.SetStrategy(strate)
	fmt.Println(com.Do())
	strate = NewStrategy("m")
	com.SetStrategy(strate)
	fmt.Println(com.Do())
	strate = NewStrategy("d")
	com.SetStrategy(strate)
	fmt.Println(com.Do())
}
