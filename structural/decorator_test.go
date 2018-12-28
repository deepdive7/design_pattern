package structural

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	ramen := Ramen{name: "ramen", price: 10}

	fmt.Println(ramen.Description())
	fmt.Println(ramen.Price())

	egg := Egg{noddles: ramen, name: "egg", price: 2}

	fmt.Println(egg.Description())
	fmt.Println(egg.Price())

	sausage := Sausage{noddles: egg, name: "sausage", price: 3}
	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())

	egg2 := Egg{noddles: egg, name: "egg", price: 2}

	fmt.Println(egg2.Description())
	fmt.Println(egg2.Price())

}
