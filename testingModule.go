package main

import "fmt"

// модуль тетстирования функций основной программы
// пишу сам для собственного потребления и тестирования частей модуля
func main() {
	//вариант для тестирования
	var variant int

	variant = 1

	if variant == 1 {
		//изначальный данные
		rx := 3
		xo := make([]int, rx*rx)
		xo[0] = 1
		xo[1] = 1
		xo[2] = 1
		xo[3] = 1
		xo[4] = 2
		xo[5] = 1
		xo[6] = 2
		xo[7] = 1
		xo[8] = 1
		var check1 int
		//выполнения функции
		flagWin := winGoriz(rx, xo, check1)
		fmt.Println(flagWin)
	}
}

//winGoriz == выполнение сравнений по горизонтали на выйгрыш
func winGoriz(rx int, xo []int, check1 int)int{
	a := 0
	for i := 0; i <= rx-1; i++ {
		for k := 0; k <= rx-1; k++ {
			if xo[a] == 1 {check1++}
			a++
			//НЕТ
			fmt.Println("a: ", a)
		}
		fmt.Println(check1)
		if check1 == rx {
			check1 = 0
			return 1
		}
		//НЕТ, показывает список
		viewXo(xo, rx)
	}
	return 0
}

func viewXo(xo []int, rx int) {
	t := 0
	for i:=0; i<=rx-1; i++ {
		for k:=0; k<=rx-1; k++ {
			fmt.Print(xo[t], " ")
			t++
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}