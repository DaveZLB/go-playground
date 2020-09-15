package factory

import "fmt"

type rectangle struct {
	Name string
}

func (r rectangle) Draw() {
	fmt.Printf("Draw %s\n",r.Name)
}