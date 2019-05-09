package method

import "fmt"

type Operation interface {
	calc(a, b int) int
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

type IFactory interface {
	getOperation() Operation
}

type AddFactory struct {
}

func (af AddFactory) getOperation() Operation {
	return Add{}
}

type SubFactory struct {
}

func (sf SubFactory) getOperation() Operation {
	return Sub{}
}

func Test() {
	factory := AddFactory{}
	fmt.Println(factory.getOperation().calc(10, 20))
}
