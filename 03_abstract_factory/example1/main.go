package main

func main() {
	chfactory := ChinaFactory{}
	appleFactory := chfactory.createAppleFactory()
	appleFactory.Produce()
	bananaFactory := chfactory.createBananaFactory()
	bananaFactory.Produce()

	jpfactory := JapanFactory{}
	appleFactory = jpfactory.createAppleFactory()
	appleFactory.Produce()
	bananaFactory = jpfactory.createBananaFactory()
	bananaFactory.Produce()
}
