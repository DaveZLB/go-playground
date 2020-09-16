package main

import "github.com/go-playground/designpatterns/abstractfactory"

func main(){
	shapeFactory := abstractfactory.GetFactory("shape")
	shapeFactory.GetShape("circle").Draw()
	shapeFactory.GetShape("rectangle").Draw()
	colorFactory := abstractfactory.GetFactory("color")
	colorFactory.GetColor("red").Fill()
	colorFactory.GetColor("green").Fill()
}