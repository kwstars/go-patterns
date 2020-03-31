package main

func main() {
	appleFactory := AppleFactory{}
	factory := appleFactory.createFactory()
	factory.Produce()

	bananaFactory := BananaFactory{}
	factory = bananaFactory.createFactory()
	factory.Produce()
}
