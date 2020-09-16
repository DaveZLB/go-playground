package color

import "fmt"

type Red struct {
	Name string
}

func (r Red) Fill() {
	fmt.Printf("Fill %s\n",r.Name)
}
