package abstract_factory

import "testing"

func TestNewCocaColaFactory(t *testing.T) {
	cocaCola := NewCocaColaFactory()
	cocaColaWater := cocaCola.CreateWater(2.5)
	cocaColaBottle := cocaCola.CreateBottle(2.5)

	cocaColaBottle.PourWater(cocaColaWater)

	if cocaColaBottle.GetWaterVolume() != cocaColaBottle.GetBottleVolume() {
		t.Errorf("Expected %v, got %v", cocaColaBottle.GetWaterVolume(), cocaColaBottle.GetBottleVolume())
	}
}
