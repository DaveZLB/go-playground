package color

import "fmt"

type Green struct {
	Name string
}

func (r Green) Fill() {
	fmt.Printf("Fill %s\n",r.Name)
}
