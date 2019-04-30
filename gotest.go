package main

import (
	"fmt"
)

func main() {
	/*
	Функция спрашивает пользователя вариант продолжения программы
	*/
	func () Varint (x string) {
		fmt.Println("1 - Добавить слово")
		fmt.Scanln(x)
		return x
	}
	var a = actionVar()
	fmt.Println(a, "выбор")

}

