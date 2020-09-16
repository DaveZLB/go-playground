package abstractfactory

import (
	"github.com/go-playground/designpatterns/abstractfactory/color"
	"github.com/go-playground/designpatterns/abstractfactory/shape"
)

type AbstractFactory interface {
	GetShape(shape string) shape.Shape
	GetColor(color string) color.Color
}