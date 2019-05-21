package decorator

import "fmt"

type Beverage interface {
	Description() string
	Cost() float64
}

type DarkRoast struct {
	description string
	cost        float64
}

func (dr *DarkRoast) Description() string {
	return dr.description
}

func (dr *DarkRoast) Cost() float64 {
	return dr.cost
}

type Mocha struct {
	description string
	cost        float64
	beverage    Beverage
}

func (m *Mocha) Description() string {
	return m.beverage.Description() + "," + m.description
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + m.cost
}

func New(beverage Beverage) *Mocha {
	return &Mocha{"12", 22, beverage}
}

func Test() {

	roast := DarkRoast{"DarkRoast", 10.0}
	mocha := &Mocha{"mocha", 22, &roast}
	fmt.Println(mocha.Description())
	fmt.Println(mocha.Cost())
}
