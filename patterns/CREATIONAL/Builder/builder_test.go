package builder

import "testing"

func TestBuilder(t *testing.T) {
	want := "<header>Header</header>" + "<body>Body</body>" + "<footer>Footer</footer>"

	product := new(Product)
	director := Director{&ConcreteBuilder{product: product}}
	director.Construct()

	got := product.Show()

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}

}
