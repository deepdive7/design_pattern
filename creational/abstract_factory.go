package creational

//AbstractFactory ...
type AbstractFactory interface {
	GetShape(string, ...float64) Shape
	GetColor(string) Color
}

// NewAbstractFactory return a specific factory obj
func NewAbstractFactory(factoryType string) AbstractFactory {
	switch factoryType {
	case "Shape":
		return NewShapeFactory()
	case "Color":
		return NewColorFactory()
	default:
		return nil
	}
}
