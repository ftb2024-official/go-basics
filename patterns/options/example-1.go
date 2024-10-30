package options

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

func NewHouse(opts ...HouseOptions) *House {
	const (
		defaultMaterial     = "wood"
		defaultHasFireplace = true
		defaultFloors       = 2
	)

	h := &House{
		defaultMaterial,
		defaultHasFireplace,
		defaultFloors,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

type HouseOptions func(h *House)

func WithConcrete() HouseOptions {
	return func(h *House) {
		h.Material = "concrete"
	}
}

func WithoutFireplace() HouseOptions {
	return func(h *House) {
		h.HasFireplace = false
	}
}

func WithFloors(floorNum int) HouseOptions {
	return func(h *House) {
		h.Floors = floorNum
	}
}

func Example1() {
	defaultHouse := NewHouse()
	newHouse := NewHouse(WithoutFireplace(), WithFloors(4))

	fmt.Println(defaultHouse) // {"wood", true, 2}
	fmt.Println(newHouse)     // {"wood", false, 4}
}
