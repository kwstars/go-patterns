package main

func main() {
	s := Subject{}
	o1 := Observer{name: "test01"}
	o2 := Observer{name: "test02"}

	s.AddObserver(&o1, &o2)

	s.NotifyObserver()
}
