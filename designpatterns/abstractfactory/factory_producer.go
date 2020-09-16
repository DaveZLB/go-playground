package abstractfactory

import "strings"

func GetFactory(factoryType string) AbstractFactory {
	if strings.ToUpper(factoryType) == "SHAPE" {
		return ShapeFactory{}
	}else if strings.ToUpper(factoryType) == "COLOR" {
		return ColorFactory{}
	}
	return nil
}


