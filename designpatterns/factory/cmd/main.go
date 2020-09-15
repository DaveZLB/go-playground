package main

import "github.com/go-playground/designpatterns/factory"

func main()  {
	f := factory.ShapeFactory{}
	f.GetShape("circle").Draw()
	f.GetShape("rectangle").Draw()
}
