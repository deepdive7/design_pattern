package structural

import "testing"

func TestFlyweight(t *testing.T) {
	CF := ColorFactory{make(map[string]Color, 10)}
	r := CF.GetColor("red")
	r.show()
	g := CF.GetColor("green")
	g.show()
	b := CF.GetColor("blue")
	b.show()
}
