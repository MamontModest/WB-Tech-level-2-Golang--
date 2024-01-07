package main

import (
	"fmt"
	"math"
)

// Плюсы: помогает определить новую операция без изменения полей самих классов,
// хорошо работает с несколькими независимыми иерархиями классов
// возможность управлять классами в одном месте
func main() {
	visit := visitor{}

	circle := &circle{3}
	triangle := &triangle{3, 4, 5}
	rectangle := &rectangle{4, 5}

	visit.rectangleSquare(rectangle)
	fmt.Println(visit.square)
	visit.circleSquare(circle)
	fmt.Println(visit.square)
	visit.triangleSquare(triangle)
	fmt.Println(visit.square)
}

// IVisitor Интерфейс visitor
type IVisitor interface {
	circleSquare(c *circle)
	triangleSquare(t *triangle)
	rectangleSquare(r *rectangle)
}

// Реализация visitor
type visitor struct {
	square float64
}

type circle struct {
	radius float64
}

func (c *circle) accept(vis IVisitor) {
	vis.circleSquare(c)
}

type triangle struct {
	//Стороны треугольника
	a, b, c float64
}

func (t *triangle) accept(vis IVisitor) {
	vis.triangleSquare(t)
}

type rectangle struct {
	width, height float64
}

func (r *rectangle) accept(vis IVisitor) {
	vis.rectangleSquare(r)
}
func (v *visitor) circleSquare(c *circle) {
	v.square = math.Pi * c.radius * c.radius
}
func (v *visitor) triangleSquare(t *triangle) {
	p := (t.a + t.b + t.c) / 2
	v.square = math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (v *visitor) rectangleSquare(r *rectangle) {
	v.square = r.height * r.width
}
