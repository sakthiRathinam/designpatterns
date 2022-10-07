package main

type Color int

type Size int

const (
	green Color = iota
	blue
	orange
)

const (
	size Size = iota
	medium
	large
)

type Product struct {
	name  string
	model string
	color Color
	size  Size
}

type Specification interface {
	IsSatisfied(prod *Product) bool
}

type ColorSpecification struct {
	filterColor Color
}
type SizeSpecification struct {
	filterSize Size
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return spec.filterColor == p.color
}
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return spec.filterSize == p.size
}

func (prod *Product) filterProd(spec int) []*Product {

}
