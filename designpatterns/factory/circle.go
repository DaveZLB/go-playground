package factory

import "fmt"

type circle struct {
	Name string
}

func (c circle) Draw() {
	fmt.Printf("Draw %s\n",c.Name)
}
