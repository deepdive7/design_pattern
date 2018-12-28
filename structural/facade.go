package structural

import (
	"fmt"
)

type zimu interface {
	showchar()
}

//A test
type A struct{}

func (a A) showchar() {
	fmt.Println("A")
}

type B struct{}

func (b B) showchar() {
	fmt.Println("B")
}

type ZiMuBiao struct {
	a A
	b B
}

func (zmb ZiMuBiao) GetA() {
	zmb.a.showchar()
}

func (zmb ZiMuBiao) GetB() {
	zmb.b.showchar()
}
