package shape

import "fmt"

type Rectangle struct {
	Name string
}

func (r Rectangle) Draw() {
	fmt.Printf("Draw %s\n",r.Name)
}