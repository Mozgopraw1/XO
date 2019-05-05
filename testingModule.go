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
		xo[4] = 1
		xo[5] = 1
		xo[6] = 1
		xo[7] = 1
		xo[8] = 2
		var check1 int
		//выполнения функции
		//flagWinG := winGoriz(rx, xo, check1)
		//fmt.Println(flagWinG)
		flagWinD1 := winDiag1(rx, xo, check1)
		fmt.Println("flagWinD1: ", flagWinD1)
		flagWinD2 := winDiag2(rx, xo, check1)
		fmt.Println("flagWinD2: ", flagWinD2)
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
			return 1
		}
		check1 = 0
		//НЕТ, показывает список
		viewXo(xo, rx)
	}
	return 0
}

//winVertic == выполнение сравнений по вертикали на выйгрыш
func winVertic(rx int, xo []int, check1 int)int{
	a := 0
	for i:= 0; i <= rx-1; i++ {
		for k :=0; k <= rx-1; k++ {
			if xo[a] == 1 {check1++}

			//НЕТ
			fmt.Println("a: ", a, "check1: ", check1)
			a = a+rx
		}
		if check1 == rx {
			return 1
		}
		check1 = 0
		a = i+1
		viewXo(xo, rx)
	}
	return 0
}

//winDiag1 == выполнение сравнения по диагонали сверху в низ слева на право
func winDiag1(rx int, xo []int, check1 int)int{
	a := 0
	for i:=0; i <=rx-1; i++ {
		if xo[a] == 1 {check1++}

		a = a+rx+1
	}
	if check1 == rx{
		return 1
	}
	check1 = 0
	viewXo(xo, rx)
	return 0
}

//winDiag2 == выполнение сравнения по диагоналис сверху в низ справа на лево
func winDiag2 (rx int, xo []int, check1 int)int{
	a := rx-1
	for i:=0; i <=rx-1; i++ {
		if xo[a] == 1 {check1++}
		a = a+rx-1
	}
	if check1 == rx {
		return 1
	}
	check1 = 0
	viewXo(xo, rx)
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