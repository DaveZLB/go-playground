package factory

import "strings"

type ShapeFactory struct {
}

func (f ShapeFactory) GetShape(shapeType string) Shape {
	if strings.ToUpper(shapeType) == "CIRCLE" {
		return circle{Name: "circle"}
	} else if strings.ToUpper(shapeType) == "RECTANGLE" {
		return rectangle{Name: "rectangle"}
	} else {
		return nil
	}
}
