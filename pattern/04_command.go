package main

import "fmt"

func main() {
	newButton := &Button{flag: false}
	fmt.Println("flag is: ", newButton.flag)
	user := UserCommand{ClickButtonOn{newButton}}

	user.execute()
	fmt.Println("flag is: ", newButton.flag)
	user = UserCommand{ClickButtonOff{newButton}}

	user.execute()
	fmt.Println("flag is: ", newButton.flag)
}

// UserCommand структура верхнего порядка
type UserCommand struct {
	Command ICommand
}

func (u *UserCommand) execute() {
	u.Command.execute()
}

type ICommand interface {
	execute()
}

type ClickButtonOn struct {
	button *Button
}

func (on ClickButtonOn) execute() {
	on.button.on()
}

type ClickButtonOff struct {
	button *Button
}

func (on ClickButtonOff) execute() {
	on.button.off()
}

type Button struct {
	flag bool
}

func (but *Button) on() {
	fmt.Println("you turn on")
	but.flag = true
}
func (but *Button) off() {
	fmt.Println("you turn off")
	but.flag = false
}
