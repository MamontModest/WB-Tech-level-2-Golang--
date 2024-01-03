package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s, err := unpack("c3b2c3def")
	fmt.Println(s, len(s), err)
}

func unpack(s string) (string, error) {
	//массив куда будем класть символы новой строки
	arr := make([]string, len(s))
	//последний символ
	lastSymbol := ""
	//число где будем хранить сколько раз нужно повторить символ
	all := 0
	//храним руны
	runes := []rune(s)
	for i, v := range runes {
		if count, err := strconv.ParseInt(string(v), 10, 64); err != nil {
			if all == 0 {
				arr = append(arr, lastSymbol)
			}
			//добавляем буквы в массив
			for i := 0; i < all; i++ {
				arr = append(arr, lastSymbol)
			}
			lastSymbol = string(v)
			all = 0
			//если буква была последней
			if i == len(runes)-1 {
				arr = append(arr, lastSymbol)
			}
		} else {
			//если число в начале строки
			if lastSymbol == "" {
				return "", errors.New("not valid string")
			}
			//если число начинается с нуля
			if count == 0 && all == 0 {
				return "", errors.New("not valid string")
			}
			//добавляем буквы в массив
			all = all*10 + int(count)
			if i == len(runes)-1 {
				for i := 0; i < all; i++ {
					arr = append(arr, lastSymbol)
				}
				all = 0
			}
		}
	}
	return strings.Join(arr, ""), nil
}
