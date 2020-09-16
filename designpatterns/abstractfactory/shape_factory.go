package abstractfactory

import (
	"github.com/go-playground/designpatterns/abstractfactory/color"
	"github.com/go-playground/designpatterns/abstractfactory/shape"
	"strings"
)

type ShapeFactory struct {
}

func (c ShapeFactory) GetShape(shapeType string) shape.Shape {
	if strings.ToUpper(shapeType) == "CIRCLE" {
		return shape.Circle{Name: "circle"}
	} else if strings.ToUpper(shapeType) == "RECTANGLE" {
		return shape.Rectangle{Name: "rectangle"}
	}
	return nil
}
func (c ShapeFactory) GetColor(colorType string) color.Color {
	return nil
}
