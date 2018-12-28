package structural

import "fmt"

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

type ColorFactory struct {
	cache map[string]Color
}

//GetColor ...
func (CF *ColorFactory) GetColor(colorType string) Color {
	if cacheColor, ok := CF.cache[colorType]; ok {
		return cacheColor
	}
	var color interface{}
	switch colorType {
	case "red":
		color = red{}
		CF.cache[colorType] = color.(Color)
	case "green":
		color = green{}
		CF.cache[colorType] = color.(Color)
	case "blue":
		color = blue{}
		CF.cache[colorType] = color.(Color)
	}
	return color.(Color)
}
