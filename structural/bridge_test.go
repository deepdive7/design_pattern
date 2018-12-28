package structural

import "testing"

func TestBridge(t *testing.T) {
	redCircle := Circle{2, 3, 4, RedCircle{}}
	redCircle.draw()
	greenCircle := Circle{6, 2, 9, GreenCircle{}}
	greenCircle.draw()
}
