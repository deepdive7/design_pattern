package creational

import (
	"errors"
	"fmt"
)

// CloneStruct ...
type CloneStruct interface {
	Clone() interface{}
}

// Person ...
type Person interface {
	CloneStruct
	Show()
}

type chinese struct {
	headColor string
	eyesColor string
	Mark      int
}

func (c chinese) Clone() interface{} {
	var pClone interface{}
	pClone = &chinese{headColor: c.headColor, eyesColor: c.eyesColor, Mark: c.Mark}
	return pClone
}

func (c chinese) Show() {
	fmt.Println("Chinese's headcolor is:", c.headColor, " EyesColor is:", c.eyesColor, ", Mark =", c.Mark)
}

type english struct {
	headColor string
	eyesColor string
	Mark      int
}

func (e english) Clone() interface{} {
	var pClone interface{}
	pClone = &english{headColor: e.headColor, eyesColor: e.eyesColor, Mark: e.Mark}
	return pClone
}

func (e english) Show() {
	fmt.Println("English's headcolor is:", e.headColor, " EyesColor is:", e.eyesColor, "Mark = ", e.Mark)
}

// ClonePerson return a person copy
func ClonePerson(personType string) (interface{}, error) {
	var p interface{}
	if _, ok := persons[personType]; ok {
		p = persons[personType].Clone()
		return p, nil
	}
	return p, errors.New("No such type person")
}
