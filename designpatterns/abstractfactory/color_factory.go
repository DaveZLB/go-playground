package abstractfactory

import (
	"github.com/go-playground/designpatterns/abstractfactory/color"
	"github.com/go-playground/designpatterns/abstractfactory/shape"
	"strings"
)

type ColorFactory struct {
}

func (c ColorFactory)GetShape(shapeType string) shape.Shape {
	return nil
}
func (c ColorFactory)GetColor(colorType string) color.Color {
	if strings.ToUpper(colorType) == "RED" {
		return color.Red{Name: "red"}
	}else if strings.ToUpper(colorType) == "GREEN" {
		return color.Green{"green"}
	}
	return nil
}
