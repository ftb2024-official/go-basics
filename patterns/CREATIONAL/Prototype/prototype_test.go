package prototype

import "testing"

func TestPrototype(t *testing.T) {
	product := NewConcreteProduct("Z")
	clonedProduct := product.Clone()

	if clonedProduct.GetName() != product.GetName() {
		t.Errorf("want product '%v', got product '%v'", product.GetName(), clonedProduct.GetName())
	}
}
