package structural

import "fmt"

//DrawAPI ...
type DrawAPI interface {
	drawCircle(int, int, int)
}

//RedCircle ...
type RedCircle struct{}

func (rc RedCircle) drawCircle(radius, x, y int) {
	fmt.Println("Drawing Circle[ color: red, radius: ", radius, ", x: ", x, ", ", y, "]")
}

//GreenCircle ...
type GreenCircle struct{}

func (gc GreenCircle) drawCircle(radius, x, y int) {
	fmt.Println("Drawing Circle[ color: green, radius: ", radius, ", x: ", x, ", ", y, "]")
}

//Shape ...
type Shape interface {
	draw()
}

//Circle ...
type Circle struct {
	x, y, radius int
	dA           DrawAPI
}

func (c Circle) draw() {
	c.dA.drawCircle(c.radius, c.x, c.y)
}
