package creational

//Packing ...
type Packing interface {
	pack() string
}

//Item ...
type Item interface {
	name() string
	packing() Packing
	price() float64
}

// Wrapper ...
type Wrapper struct{}

func (w Wrapper) pack() string {
	return "Wrapper"
}

//Bottle ...
type Bottle struct{}

func (b Bottle) pack() string {
	return "Bottle"
}

//VegBurger ...
type VegBurger struct {
	Bottle
}

func (vb VegBurger) name() string {
	return "Veg Burger"
}

func (vb VegBurger) price() float64 {
	return 12.0
}

// ChickBurger ...
type ChickBurger struct {
	Wrapper
}

func (cb ChickBurger) name() string {
	return "Chick Burger"
}

func (cb ChickBurger) price() float64 {
	return 24.0
}
