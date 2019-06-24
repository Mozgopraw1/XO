package main

import "fmt"

// модуль тетстирования функций основной программы
// пишу сам для собственного потребления и тестирования частей модуля

type strData struct {
	rx int
	check1 int
}

func main() {
	str := new(strData)

	//вариант для тестирования
	var variant int

	variant = 1

	if variant == 1 {

		//изначальный данные
		str.rx = 3
		xo := make([]int, str.rx*str.rx)
		xo[0] = 0
		xo[1] = 1
		xo[2] = 0
		xo[3] = 0
		xo[4] = 1
		xo[5] = 0
		xo[6] = 0
		xo[7] = 1
		xo[8] = 0
		//выполнения функции
		//flagWinG := winGoriz(rx, xo, check1)
		//fmt.Println(flagWinG)
		viewXO(xo, *str)
		flagWinD1 := winDiag1(*str, xo)
		fmt.Println("flagWinD1: ", flagWinD1)
		flagWinD2 := winDiag2(*str, xo)
		fmt.Println("flagWinD2: ", flagWinD2)
		flagWinV := winVertic(*str, xo)
		fmt.Println("flagWinV: ", flagWinV)
		flagWinG := winGoriz(*str, xo)
		fmt.Println("flagWinG: ", flagWinG)
	}
}

//Выйгрыш
func winX(str strData, xo []int){
	viewXO(xo, str)
	flagWinD1 := winDiag1(str, xo)
	fmt.Println("flagWinD1: ", flagWinD1)
	flagWinD2 := winDiag2(str, xo)
	fmt.Println("flagWinD2: ", flagWinD2)
	flagWinV := winVertic(str, xo)
	fmt.Println("flagWinV: ", flagWinV)
	flagWinG := winGoriz(str, xo)
	fmt.Println("flagWinG: ", flagWinG)
}
//winGoriz == выполнение сравнений по горизонтали на выйгрыш
func winGoriz(str strData, xo []int)int{
	a := 0
	for i := 0; i <= str.rx-1; i++ {
		for k := 0; k <= str.rx-1; k++ {
			if xo[a] == 1 {str.check1++}
			a++
		}
		if str.check1 == str.rx {
			return 1
		}
		str.check1 = 0
	}
	return 0
}

//winVertic == выполнение сравнений по вертикали на выйгрыш
func winVertic(str strData, xo []int)int{
	a := 0
	for i:= 0; i <= str.rx-1; i++ {
		for k :=0; k <= str.rx-1; k++ {
			if xo[a] == 1 {str.check1++}
			a = a+str.rx
		}
		if str.check1 == str.rx {
			return 1
		}
		str.check1 = 0
		a = i+1
	}
	return 0
}

//winDiag1 == выполнение сравнения по диагонали сверху в низ слева на право
func winDiag1(str strData, xo []int)int{
	a := 0
	for i:=0; i <=str.rx-1; i++ {
		if xo[a] == 1 {str.check1++}

		a = a+str.rx+1
	}
	if str.check1 == str.rx{
		return 1
	}
	str.check1 = 0
	return 0
}

//winDiag2 == выполнение сравнения по диагонали сверху в низ справа на лево
func winDiag2 (str strData, xo []int)int{
	a := str.rx-1
	for i:=0; i <=str.rx-1; i++ {
		if xo[a] == 1 {str.check1++}
		a = a+str.rx-1
	}
	if str.check1 == str.rx {
		return 1
	}
	str.check1 = 0
	return 0
}


func viewXO(xo []int, str strData) {
	t := 0
	for i:=0; i<=str.rx-1; i++ {
		for k:=0; k<=str.rx-1; k++ {
			fmt.Print(xo[t], " ")
			t++
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}