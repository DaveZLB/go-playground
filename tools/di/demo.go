package main

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Greeter struct {
	Message Message
	Grumpy  bool
}

type Event struct {
	Greeter Greeter
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)

}

//普通方式
func main() {
	message := NewMessage("Hi,I am here")
	greeter := NewGreeter(message)
	event,err := NewEvent(greeter)
	if err != nil {
		fmt.Printf("%v\n",err)
		return
	}

	event.Start()
}
