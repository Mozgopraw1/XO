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
		winGoriz(rx, xo, check1) }
}

//winGoriz == выполнение сравнений по горизонтали на выйгрыш
func winGoriz(rx int, xo []int, check1 int){
	a := 0
	for i := 0; i <= rx-1; i++ {
		for k := 0; k <= rx-1; k++ {
			if xo[a] == 1 {check1++}
			a++
		}
		//НЕТ, показывает список
		for l:=1; l<=(rx*rx)-1; l++ {fmt.Print(xo[l])}

		 if check1 == rx-1 {fmt.Println("равно check1 = ", check1, ", rx-1 = ", rx-1)}
		check1 = 0
	}
}