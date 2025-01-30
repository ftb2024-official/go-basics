package factory_method

import "log"

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

// Creator provides a factory interface
type Creator interface {
	CreateProduct(action action) Product // factory method
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

type Product interface {
	Use() string // every product should be usable
}

type ConcreteCreator struct{}

func (cc *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ProductA{action: string(action)}
	case B:
		product = &ProductB{action: string(action)}
	case C:
		product = &ProductC{action: string(action)}
	default:
		log.Fatalln("Unknown action")
	}

	return product
}

type ProductA struct {
	action string
}

func (a *ProductA) Use() string {
	return a.action
}

type ProductB struct {
	action string
}

func (b *ProductB) Use() string {
	return b.action
}

type ProductC struct {
	action string
}

func (c *ProductC) Use() string {
	return c.action
}
