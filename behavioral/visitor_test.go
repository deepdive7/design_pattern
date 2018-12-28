package behavioral

import "testing"

func TestVisitor(t *testing.T) {
	Data := &ABData{A: 8, B: 10}
	add := &AddVisitor{}
	sub := &SubVisitor{}
	Data.Accept(add)
	Data.Accept(sub)
}
