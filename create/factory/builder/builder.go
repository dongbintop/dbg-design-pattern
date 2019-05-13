package builder

import "fmt"

type Parlour struct {
	wall string
	tv   string
	sofa string
}

type Builder interface {
	build() Parlour
	wall(string) Builder
	tv(string) Builder
	sofa(string) Builder
}

type ParlourBuilder struct {
	parlour Parlour
}

func (pb *ParlourBuilder) wall(wall string) Builder {
	pb.parlour.wall = wall
	return pb
}

func (pb *ParlourBuilder) tv(tv string) Builder {
	pb.parlour.tv = tv
	return pb
}

func (pb *ParlourBuilder) sofa(sofa string) Builder {
	pb.parlour.sofa = sofa
	return pb
}

func (pb *ParlourBuilder) build() Parlour {
	return pb.parlour
}

func Test() {
	builder := &ParlourBuilder{parlour: Parlour{}}
	p := builder.wall("bai").tv("xioami").sofa("geli").build()
	fmt.Println(p)
}
