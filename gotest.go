package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ход игрока; true == x; false o == false
	flag := true

	// доступен ли ход в точку
	flagIf := false

	//Включает рандомную функцию
	rand.Seed(time.Now().UTC().UnixNano())

	// создание массива
	xo := make([]int, 9)

	// значение точки для хода
	var x int

	//запуск всей игры
	xod(xo, x, flag, flagIf)
}

//основная функция по запуску игры
func xod(xo []int, x int, flag bool, flagIf bool) {
	for i := 0; i <= 8; i++ {
		x = rand.Intn(8)            // рандомная цифра от 0 до 8
		vievXo(xo, flag, flagIf, x)    // показ поля
		xoVariant(flag, x, xo, flagIf) // проверка хода
		xoStageAssign(xo, x, flag)     // функция присвоения

		flag1(&flag) 					// смена стороны
	}
}

//вывод поля
func vievXo(xo []int, flag bool, flagIf bool, x int) {
	fmt.Println(xo[0], xo[1], xo[2])
	fmt.Println(xo[3], xo[4], xo[5])
	fmt.Println(xo[6], xo[7], xo[8])
	fmt.Println(" ")

	fmt.Println(flag, flagIf, x)
}

//смена хода
func flag1(flag *bool) {
	if *flag{
		*flag = false
	} else {
		*flag = true
	}
}

// ход на поле
// проверка чей ход
func xoVariant(flag bool, x int, xo []int, flagIf bool) {
	// X ход
	if flag{
		xoCor(xo, &x, &flagIf)
		flagIf = false
	}
	// O ход
	if flag{
		xoCor(xo, &x, &flagIf)
		flagIf = false
	}
}

// проверка был ли уже ход в данной точке
// я запутался в указателях именно тут, все тесты показывают что X - как точка куда ставится X или O
// не меняется вне функции и в последующей функции, но в самой функции вроде как меняется (вроде даже в xoCorDop тоже)
func xoCor(xo []int, x *int, flagIf *bool) {
	if xo[*x] == 0 {
		*flagIf = true
	}
	if !*flagIf {
		xoCorDop(xo, x, flagIf)
	}
}

// если точка уже занята
func xoCorDop(xo []int, x *int, flagIf *bool) {
	*flagIf = false
	*x = rand.Intn(8) // сразу даёт рандомное значение
	xoCor(xo, x, flagIf)
}

//функция присвоения значения проверки
func xoStageAssign(xo []int, x int, flag bool) {
	if flag{ //ход X
		xo[x] = 1
	}
	if !flag{ //ход O
		xo[x] = 2
	}
}
