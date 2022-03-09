package main

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

type Rice struct {
	Product
}
type Pasta struct {
	Product
}

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NameVisitor struct {
	ProductList []string
}

func (nv *NameVisitor) Visit(p ProductInfoRetriever) {
	nv.ProductList = append(nv.ProductList, p.GetName())
}

func main() {
	products := make([]Visitable, 3)
	products[0] = &Rice{Product{32.00, "Rice"}}
	products[1] = &Pasta{Product{19.00, "Pasta"}}
	products[2] = &Fridge{Product{333.00, "Fridge"}}

	priceVisitor := &PriceVisitor{}
	nameVisitor := &NameVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
		p.Accept(nameVisitor)
	}
	fmt.Printf("Total: %f\n", priceVisitor.Sum)
	fmt.Printf("Name: %v\n", nameVisitor.ProductList)

}
