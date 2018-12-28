package creational

import (
	"fmt"
	"testing"
)

func TestShapeFactory(t *testing.T) {
	sf := NewAbstractFactory("Shape")
	if sf == nil {
		return
	}
	radius := 1.0
	c := sf.GetShape("circle", radius)
	fmt.Printf("circle(%.02f)'s area = %.02f\n", radius, c.area())
}

func TestColorFactory(t *testing.T) {
	cf := NewAbstractFactory("Color")
	if cf == nil {
		return
	}
	r := cf.GetColor("red")
	r.show()
}
