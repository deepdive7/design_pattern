package creational

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	fmt.Println("Buy a Veg Burger")
	vb := &VegBurger{}
	fmt.Printf("%s(%s)'s price is %.02f\n", vb.name(), vb.pack(), vb.price())
	fmt.Println("Buy a Chick Burger")
	cb := &ChickBurger{}
	fmt.Printf("%s(%s)'s price is %.02f\n", cb.name(), cb.pack(), cb.price())
}
