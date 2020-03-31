package main

type Factory interface {
	createAppleFactory() Product
	createBananaFactory() Product
}

type ChinaFactory struct{}

func (a *ChinaFactory) createAppleFactory() Product {
	return new(ChinaApple)
}

func (a *ChinaFactory) createBananaFactory() Product {
	return new(ChinaBanana)
}

type JapanFactory struct{}

func (a *JapanFactory) createAppleFactory() Product {
	return new(JapanApple)
}

func (a *JapanFactory) createBananaFactory() Product {
	return new(JapanBanana)
}
