package main

import (
	"errors"
	"fmt"
)

// Плюсы: единая точка входа, объект скрывает бизнес логику и предоставляет лишь интерфейс пользователю, может делегировать задачи

//Минусы: В какой-то момент объект может очень сильно разрастись и им будет тяжело управлять и модифицировать.
//Читабельность кода

// Реальное использование: например при написании API сервера, для клиента установленна только одна точка входа, запрос
// к http
func main1() {
	window := Window{
		iMarket: Market{Address: "Москва ул Пушкина 127 п.5 кв98"},
		iCash:   Cash{Money: 1500},
		iUser: User{
			Name:    "Vania",
			Surname: "Popov",
		},
	}
	window.FirstDelivery()
}

// Window facade объединённая точка входа в наше приложение, которая совмещает в себе 3 интерфейса.
// Делегирует вызовы функций нижестоящим интерфейсам и скрывает для пользователя сложность архитектуры
type Window struct {
	iUser
	iMarket
	iCash
}

// FirstDelivery пользователь делает первую доставку
func (window Window) FirstDelivery() {
	window.greeting()
	err := window.billing(1000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	window.orderDelivery()
}

type iCash interface {
	billing(price int) error
}

type Cash struct {
	Money int
}

func (c Cash) billing(price int) error {
	if c.Money < price {
		return errors.New("not enough money\n")
	}
	fmt.Printf("you pay %d\n", price)
	return nil
}

// interface market
type iMarket interface {
	orderDelivery()
}

// Market магазин по доставке продуктов
type Market struct {
	Address string
}

func (m Market) orderDelivery() {
	fmt.Printf("Products have been ordered to be delivered to your address: %s\n", m.Address)
}

// interface user
type iUser interface {
	greeting()
}

// User пользователь системы
type User struct {
	Name,
	Surname string
}

//приветствие в главном меню

func (u User) greeting() {
	fmt.Printf("hello my dear %s %s\n", u.Name, u.Surname)
}
