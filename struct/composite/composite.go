package composite

import "fmt"

type MenuDesc struct {
	name        string
	description string
}

func (md *MenuDesc) Name() string {
	return md.name
}

func (md *MenuDesc) Description() string {
	return md.description
}

type Group interface {
	Add(component MenuComponent)
	Remove(idx int)
	Child(idx int) MenuComponent
}

//type Product interface {
//
//}

type MenuComponent interface {
	Print()
	Price() float32
}

type MenuItem struct {
	MenuDesc
	price float32
}

func (mi *MenuItem) Price() float32 {
	return mi.price
}

func (mi *MenuItem) Print() {
	fmt.Printf("  %s, ￥%.2f\n", mi.name, mi.price)
	fmt.Printf("    -- %s\n", mi.description)
}

func NewMenuItem(name, description string, price float32) *MenuItem {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
		price: price,
	}
}

type MenuGroup struct {
	children []MenuComponent
}

type Menu struct {
	MenuDesc
	MenuGroup
}


func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ￥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
	fmt.Println()
}

func (m *MenuGroup) Add(c MenuComponent) {
	m.children = append(m.children, c)
}

func (m *MenuGroup) Remove(idx int) {
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}

func (m *MenuGroup) Child(idx int) MenuComponent {
	return m.children[idx]
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}

func Test()  {
	menu1 := NewMenu("培根鸡腿燕麦堡套餐", "供应时间：09:15--22:44")
	menu1.Add(NewMenuItem("主食", "培根鸡腿燕麦堡1个", 11.5))
	menu1.Add(NewMenuItem("小吃", "玉米沙拉1份", 5.0))
	menu1.Add(NewMenuItem("饮料", "九珍果汁饮料1杯", 6.5))

	menu2 := NewMenu("奥尔良烤鸡腿饭套餐", "供应时间：09:15--22:44")
	menu2.Add(NewMenuItem("主食", "新奥尔良烤鸡腿饭1份", 15.0))
	menu2.Add(NewMenuItem("小吃", "新奥尔良烤翅2块", 11.0))
	menu2.Add(NewMenuItem("饮料", "芙蓉荟蔬汤1份", 4.5))

	menu3 := NewMenu("台湾卤肉饭套餐", "供应时间：09:15--22:44")
	menu3.Add(NewMenuItem("主食", "台湾卤肉饭1份", 22.0))
	menu3.Add(NewMenuItem("小吃", "甜甜圈", 4.0))
	menu3.Add(NewMenuItem("饮料", "可乐1杯", 6.5))

	all := NewMenu("超值午餐", "周一至周五有售")
	all.Add(menu1)
	all.Add(menu2)
	all.Add(menu3)

	all.Print()
}

