package creational

import "testing"

func TestPrototype(t *testing.T) {
	var c1, c2 *chinese
	var e1, e2 *english
	if ci, err := ClonePerson("chinese"); err == nil {
		c1 = ci.(*chinese)
	}
	if ci, err := ClonePerson("chinese"); err == nil {
		c2 = ci.(*chinese)
	}
	if ei, err := ClonePerson("english"); err == nil {
		e1 = ei.(*english)
	}
	if ei, err := ClonePerson("english"); err == nil {
		e2 = ei.(*english)
	}
	c1.Mark = 1
	c2.Mark = 2
	e1.Mark = 3
	e2.Mark = 4
	c1.Show()
	c2.Show()
	e1.Show()
	e2.Show()
}
