package main

type Factory interface {
	createFactory() Product
}

type AppleFactory struct{}

func (a *AppleFactory) createFactory() Product {
	return new(Apple)
}

type BananaFactory struct{}

func (a *BananaFactory) createFactory() Product {
	return new(Banana)
}
