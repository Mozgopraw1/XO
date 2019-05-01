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
		x = rand.Intn(8) // рандомная цифра от 0 до 8

		xoVariant(x, xo, flagIf, flag)  // проверка хода
		//xoStageAssign(&xo, x, flag)   // функция присвоения
		vievXo(xo, flag, flagIf, x)     // показ поля
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
func xoVariant(x int, xo []int, flagIf bool, flag bool) {
		xoCor(xo, x, &flagIf, flag)
		flagIf = false
}

// проверка был ли уже ход в данной точке
// я запутался в указателях именно тут, все тесты показывают что X - как точка куда ставится X или O
// не меняется вне функции и в последующей функции, но в самой функции вроде как меняется (вроде даже в xoCorDop тоже)
func xoCor(xo []int, x int, flagIf *bool, flag bool) {
	for i:=0; i<=2; i++ {
		fmt.Println("тест 1", x)
		if xo[x] == 0 {
			*flagIf = true
			if flag{ //ход X
				xo[x] = 1
			}
			if !flag{ //ход O
				xo[x] = 2
			}
			i = 3
		}
		if !*flagIf {
			x = rand.Intn(9) // сразу даёт рандомное значение
			fmt.Println("тест 2", x)
			i = 0
		}
	}
}



//функция присвоения значения проверки
/*func xoStageAssign(xo *[]int, x *int, flag bool) {
	if flag{ //ход X
		xo[x] = 1
	}
	if !flag{ //ход O
		xo[*x] = 2
	}
}
*/