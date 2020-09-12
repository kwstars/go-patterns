package main

type ISubject interface {
	AddObserver(observer ...IObserver)
	//RemoveObserver(observer IObserver)
	NotifyObserver()
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
