package abstract_factory

// an interface for creating families of related objects
type AbstractFactory interface {
	CreateWater(volume float64) AbstractWater
	CreateBottle(volume float64) AbstractBottle
}

type AbstractWater interface {
	GetVolume() float64
}

type AbstractBottle interface {
	PourWater(water AbstractWater)
	GetBottleVolume() float64
	GetWaterVolume() float64
}

type CocaColaFactory struct{}

func NewCocaColaFactory() AbstractFactory {
	return &CocaColaFactory{}
}

func (f *CocaColaFactory) CreateWater(volume float64) AbstractWater {
	return &CocaColaWater{volume: volume}
}

func (f *CocaColaFactory) CreateBottle(volume float64) AbstractBottle {
	return &CocaColaBottle{volume: volume}
}

type CocaColaWater struct {
	volume float64
}

// GetVolume returns volume of drink
func (w *CocaColaWater) GetVolume() float64 {
	return w.volume
}

// CocaColaBottle implements AbstractBottle
type CocaColaBottle struct {
	water  AbstractWater
	volume float64
}

// PourWater pours water into bottle
func (b *CocaColaBottle) PourWater(water AbstractWater) {
	b.water = water
}

// GetBottleVolume returns volume of bottle
func (b *CocaColaBottle) GetBottleVolume() float64 {
	return b.volume
}

// GetWaterVolume returns volume of water
func (b *CocaColaBottle) GetWaterVolume() float64 {
	return b.water.GetVolume()
}
