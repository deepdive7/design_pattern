package creational

import (
	"fmt"
	"math"
)

// ShapeFactory is used to generate diff shapes
type ShapeFactory struct{}

// GetShape generates a specific Shape obj
func (SF ShapeFactory) GetShape(shapeType string, args ...float64) Shape {
	switch shapeType {
	case "circle":
		return circle{args[0]}
	case "square":
		return square{args[0]}
	case "rectangle":
		return rectangle{args[0], args[1]}
	default:
		return nil
	}
}

// GetColor is a empty method to impl abstractFactory
func (SF ShapeFactory) GetColor(colorType string) Color {
	return nil
}

// NewShapeFactory return a shapeFactory obj
func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{}
}

//Shape ...
type Shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return s.sideLen * s.sideLen
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

// ColorFactory generate color obj
type ColorFactory struct{}

// GetShape is a empty method to impl abstractFactory
func (CF ColorFactory) GetShape(shapeType string, args ...float64) Shape {
	return nil
}

// GetColor returns a specific a color obj
func (CF ColorFactory) GetColor(colorType string) Color {
	switch colorType {
	case "red":
		return &red{}
	case "green":
		return &green{}
	case "blue":
		return &blue{}
	default:
		return nil
	}
}

//NewColorFactory return a new ColorFactory obj
func NewColorFactory() *ColorFactory {
	return &ColorFactory{}
}

//Color ...
type Color interface {
	show()
}

type red struct{}

func (r red) show() {
	fmt.Println("I am red")
}

type green struct{}

func (g green) show() {
	fmt.Println("I am green")
}

type blue struct{}

func (b blue) show() {
	fmt.Println("I am blue")
}
