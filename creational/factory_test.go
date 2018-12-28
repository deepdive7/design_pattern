package creational

import (
	"fmt"
	"testing"
)

func TestCircle(t *testing.T) {
	radius := 1.0
	c := SF.GetShape("circle", radius)
	if c == nil {
		return
	}
	fmt.Printf("circle(%.02f)'s area = %.02f\n", radius, c.area())
}

func TestSquare(t *testing.T) {
	sideLen := 10.0
	s := SF.GetShape("square", sideLen)
	if s == nil {
		return
	}
	fmt.Printf("square(%.02f)'s area = %.02f\n", sideLen, s.area())
}
func TestRectangle(t *testing.T) {
	height := 10.0
	width := 5.0
	r := SF.GetShape("rectangle", height, width)
	if r == nil {
		return
	}
	fmt.Printf("rectangle(%.02f, %.02f)'s area = %.02f\n", height, width, r.area())
}

func TestRed(t *testing.T) {
	r := CF.GetColor("red")
	r.show()
}

func TestGreen(t *testing.T) {
	g := CF.GetColor("green")
	g.show()
}

func TestBlue(t *testing.T) {
	b := CF.GetColor("blue")
	b.show()
}
