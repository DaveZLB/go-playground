package main

import (
	"errors"
	"fmt"
	"go.uber.org/dig"
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

func main() {
	c := dig.New()
	_ = c.Provide(func() string {
		return "Hi,Tom"
	})
	_ = c.Provide(NewMessage)
	_ = c.Provide(NewGreeter)
	_ = c.Provide(NewEvent)
	_ = c.Invoke(func(e Event) {
		e.Start()
	})
}
