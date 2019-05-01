package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// ход игрока; true == x; false o == false
	flag := true

	// доступен ли ход в точку
	flagIf := false

	var game int
	newGame(&game)


	//Включает рандомную функцию
	rand.Seed(time.Now().UTC().UnixNano())

	// создание массива
	xo := make([]int, 9)

	// значение точки для хода
	var x int
	// 0 ни кто не выйграл, 1 выйграл Х, 2 выйграл О
	var flagWin int

	//сколько побед
	Xwin := 0
	Owin := 0
	NoWin := 0

	//запуск всей игры
	for i:=1; i<=game; i++ {
		xo = make([]int, 9)
		xod(xo, x, flag, flagIf, &flagWin, &NoWin, &Xwin, &Owin)
	}

	fmt.Println("Общий результат:")
	fmt.Println("X:     ", Xwin)
	fmt.Println("O:     ", Owin)
	fmt.Println("Ничья: ", NoWin)
	}

//основная функция по запуску игры
func xod(xo []int, x int, flag bool, flagIf bool, flagWin *int, NoWin *int, Xwin *int, Owin *int) {
	for i := 0; i <= 8; i++ {
		x = rand.Intn(8) // рандомная цифра от 0 до 8

		xoVariant(x, xo, flagIf, flag)  // проверка хода
		vievXo(xo)     // показ поля
		win(xo, flagWin, Xwin, Owin)
		flag1(&flag) 					// смена стороны
		if *flagWin != 0 {
			fmt.Println("Действительно выйграл, выходим")
			i = 9
			*flagWin = 0 //сброс значения
		}
		if *flagWin == 0 && i==8 { fmt.Println("Похоже ни кто не выйграл")
		*NoWin++} //проверка на ничью
	}
}

//вывод поля
func vievXo(xo []int) {
	fmt.Println(xo[0], xo[1], xo[2])
	fmt.Println(xo[3], xo[4], xo[5])
	fmt.Println(xo[6], xo[7], xo[8])
	fmt.Println(" ")
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
func xoCor(xo []int, x int, flagIf *bool, flag bool) {
	for i:=0; i<=2; i++ {
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
			i = 0
		}
	}
}

//проверка совпадений линий
func win(xo []int, flagWin *int, Xwin *int, Owin *int) {
	for i:=1; i<=2; i++ {
		if xo[0]&xo [1]&xo[2] == i {
			*flagWin = i
		}
		if xo[3]&xo [4]&xo[5] == i {
			*flagWin = i
		}
		if xo[6]&xo [7]&xo[8] == i {
			*flagWin = i
		}
		if xo[0]&xo [3]&xo[6] == i {
			*flagWin = i
		}
		if xo[1]&xo [4]&xo[7] == i {
			*flagWin = i
		}
		if xo[2]&xo [5]&xo[8] == i {
			*flagWin = i
		}
		if xo[0]&xo [4]&xo[8] == i {
			*flagWin = i
		}
		if xo[2]&xo [4]&xo[6] == i {
			*flagWin = i
		}
	}
	if *flagWin == 1 {
		fmt.Println("Победа X")
		*Xwin++
	}
	if *flagWin == 2 {
		fmt.Println("Победа O")
		*Owin++
	}
}

func newGame(game *int) {
	fmt.Println("Приветствую тебя в игре XO, сколько желаешь сыграть партий?")
	var st string

	fmt.Fscan(os.Stdin, &st)

	i1, err := strconv.Atoi(st)
	if err == nil {
		*game = i1
	}
}