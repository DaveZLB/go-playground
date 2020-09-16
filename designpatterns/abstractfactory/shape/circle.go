package shape

import "fmt"

type Circle struct {
	Name string
}

func (c Circle) Draw() {
	fmt.Printf("Draw %s\n",c.Name)
}
