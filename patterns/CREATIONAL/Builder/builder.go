package builder

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

type ConcreteBuilder struct {
	product *Product
}

func (cb *ConcreteBuilder) MakeHeader(str string) {
	cb.product.Content += "<header>" + str + "</header>"
}

func (cb *ConcreteBuilder) MakeBody(str string) {
	cb.product.Content += "<body>" + str + "</body>"
}

func (cb *ConcreteBuilder) MakeFooter(str string) {
	cb.product.Content += "<footer>" + str + "</footer>"
}

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}
