package main

import "fmt"

type IObserver interface {
	Update()
}

type ISubject interface {
	AddObserver(observer ...IObserver)
	//RemoveObserver(observer IObserver)
	NotifyObserver()
}

type Observer struct {
	name string
}

func (o *Observer) Update() {
	fmt.Printf("%s触发\n", o.name)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObserver(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject) NotifyObserver() {
	for i := range s.observers {
		s.observers[i].Update()
	}
}

func main() {
	s := new(Subject)
	o1 := Observer{name: "test01"}
	o2 := Observer{name: "test02"}

	s.AddObserver(&o1)
	s.AddObserver(&o2)

	s.NotifyObserver()
}
