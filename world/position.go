package world

import "fmt"

type Position struct {
	x int
	y int
}

func NewPosition() *Position {
	return &Position{
		x: 0,
		y: 0,
	}
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p *Position) Up() {
	p.y++
}

func (p *Position) Down() {
	p.y--
}

func (p *Position) Left() {
	p.x--
}

func (p *Position) Right() {
	p.x++
}

func (p *Position) X() int {
	return p.x
}

func (p *Position) Y() int {
	return p.y
}
