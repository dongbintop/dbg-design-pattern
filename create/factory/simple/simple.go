package simple

import "fmt"

type Operation interface {
	calc(a int, b int) int
}

type Add struct {
}

func (add Add) calc(a, b int) int {
	return a + b
}

type Sub struct {
}

func (sub Sub) calc(a, b int) int {
	return a - b
}

type Mul struct {
}

func (mul Mul) calc(a, b int) int {
	return a * b
}

type Div struct {
}

func (div Div) calc(a, b int) int {
	return a / b
}

type Factory struct {
}

func (s Factory) createOperation(t string) Operation {
	switch t {
	case "-":
		return Sub{}
	case "*":
		return Mul{}
	case "/":
		return Div{}
	default:
		return Add{}
	}
}

func Test() {
	factory := Factory{}
	op := factory.createOperation("*")
	fmt.Println(op.calc(10, 12))
}
