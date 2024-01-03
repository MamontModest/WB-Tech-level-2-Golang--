package main

import "fmt"

// Абстрактный интерфейс для создания объектов (продукта).
// плюсы: при больших структурах упрощает и делает более понятным процесс создание различных экземпляров класса
// так-же делегирует создание объекта, отделяет построение сложного объекта от его представления

// Минусы классы должны быть изменяемыми, для каждого класса нужен свой builder (много однотипного кода)
func main2() {
	d := DinningRoom{lunch{}}
	d.goLunch()
}

// DinningRoom Место где кушают, в данном случае является директором builder
type DinningRoom struct {
	lunch
}

// Создаём какой нам нужен обед
func (d DinningRoom) goLunch() {
	d.moreMeat(150)
	d.moreRice(100)
	d.moreSauce(30)
	d.moreVegetables(179)
	d.setUpDrink("soda")
	d.getLunch()
}

// Сложный объект, который нуждается в гибкости построения
type lunch struct {
	vegetables int
	meat       int
	rice       int
	sauce      int
	drink      string
}

// методы для управления полями lunch

func (l *lunch) moreVegetables(grams int) {
	l.vegetables += grams
}
func (l *lunch) moreMeat(grams int) {
	l.meat += grams
}

func (l *lunch) moreSauce(grams int) {
	l.sauce += grams
}

func (l *lunch) moreRice(grams int) {
	l.sauce += grams
}
func (l *lunch) setUpDrink(d string) {
	l.drink += d
}

// Результат собранного нами ланча
func (l *lunch) getLunch() {
	fmt.Printf(
		"today you will be eat:\n%d gramms meat\n%d gramms rice\n%d gramms vegatables\n%d gramms sauce\nand drink %s",
		l.meat, l.rice, l.vegetables, l.sauce, l.drink,
	)
}
