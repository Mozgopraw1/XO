package main

import (
	"fmt"
	"math/rand"
	"time"
)

// REVIEW: зачем ты вообще комментируешь неэкспортированные
//  функции? Хотя, это не запрещено, но зачем? И если
//  комментируешь, делай это правильно!

// Понял, буду пытаться коментировать как нужно

// REVIEW: коментарии в коде должны комментировать что-то
//  действительно неочевидное. Комментировать простейшие
//  вещи не принято. Это описано в статье про психологию
//  простого кода.

// REVIEW: есть переменные, которые проваливаются слишком
//  глубоко в стек вызовов, это создаёт разрыв контекста и
//  сильно усложняет чтение.

// Хорошо

// структура данных.
type strData struct {
	rx int		// == ширина/длина поля
	check1 int	// == временная перемення в winD,V,G
	unit int 	// == крестик или нолик (1 или 2)
	winF int 	// == win передача флага об проверке
	flagIf bool // == флаг свободной точки
	flagWin int // == 0 ни кто не выйграл, 1 выйграл Х, 2 выйграл О
	xWin int 	// == количество побед X
	oWin int 	// == количество побед O
	noWin int 	// == количество ничьих
	flag bool 	// == ход игрока; true == x; false == o;
	game int	// == количество партий
	err error  // == ошибки, для обработки.
}

func main() {
	//инициализация структыр
	str := new(strData)
	str.flagIf = false

	// ширина/длина поля
	str.rx = 3
	// количество партий которые хочет сыграть игрок
	str.game = 100

	//Включает рандомную функцию
	rand.Seed(time.Now().UTC().UnixNano())

	//сколько побед
	str.xWin = 0
	str.oWin = 0
	str.noWin = 0

	//ввод основных данных о игре
	str.err = newGame(str)

	//узнаёт время в данный момент
	dt := time.Now()

	//запуск всей игры
	for i := 1; i <= str.game; i++ {
		xo := make([]int, str.rx*str.rx)
		xod(xo, str)
	}

	// общая статистика в конце игры
	fmt.Println("Общий результат:")
	fmt.Println("X:     ", str.xWin)
	fmt.Println("O:     ", str.oWin)
	fmt.Println("Ничья: ", str.noWin)
	dt1 := time.Now()
	fmt.Println("Затраченное время на вычисление: ", dt1.Sub(dt))
}

//xod == основная функция по запуску игры(буду переделывать, заменяя набор переменных на структуры)
func xod(xo []int, str *strData) {
	for i := 0; i <= str.rx*str.rx-1; i++ {

		//X == координата точки на поле
		x := rand.Intn(str.rx*str.rx-1) // рандомная цифра от 0 до RX*RX
		xoVariant(x, xo, *str) 		// проверка хода
		str.winF = win(xo, str)     // проверка по вертикали, горизонтали, диагонали на выйгрышь
		if str.winF >= 1 {
			i = str.rx*str.rx+1
			viewXo(xo, *str)        // показ поля
		}
		str.flag = playerMove(*str) // смена стороны

		// проверка на ничью, и сброс флага выйгрыша для начала новой партии
		if str.flagWin != 0 {
			i = str.rx*str.rx+1
		}
		if str.flagWin == 0 && i == str.rx*str.rx-1 {

			fmt.Println("Похоже ни кто не выйграл")
			viewXo(xo, *str)        // показ поля
			str.noWin++
		}
		str.flagWin = 0 			//сброс значения
	}
}

//newGame == ввод основных данных об игре
func newGame (str *strData)error{
	fmt.Println("Привет, какой ширины/длины ты хочешь поле?")
	_, err := fmt.Scan(&str.rx)
	if err != nil {
		return err
	}
	fmt.Println("А сколько партий ты желаешь сыграть?")
	_, err = fmt.Scan(&str.game)
	if err != nil {
		return err
	}
	return str.err
}

//viewXo == вывод поля в консоль
func viewXo(xo []int, str strData) {
	t := 0
	for i := 0; i <= str.rx-1; i++ {
		for k := 0; k <= str.rx-1; k++ {
			viewXoEnd(xo, t)
			t++
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

//vievXoEnd == замена 1 и 2 на X и O
func viewXoEnd(xo []int, t int) {
	if xo[t] == 0 { fmt.Print("_ ")}
	if xo[t] == 1 { fmt.Print("X ")}
	if xo[t] == 2 { fmt.Print("O ")}
}

//playerMove == смена хода
func playerMove(str strData)bool{
	if str.flag {
		return false
	} else {
		return true
	}
}

// ход на поле
// xoVariant == проверка чей ход
func xoVariant(x int, xo []int, str strData) {
	xoCor(xo, x, str)
	str.flagIf = false
}

// xoCor == проверка был ли уже ход в данной точке
// REVIEW: высокая цикломатика (6, надо 5 и меньше)
// не знаю как цикломатику уменьшить в данном случае, описал подробно каждую строку кода тут
// если разделять данную функцию, то я уже потеряю смысл её.
func xoCor(xo []int, x int, str strData) {
	for i := 0; i <= 2; i++ { // повторяет цикл бесконечно пока точка на поле не будет пустой
		if xo[x] == 0 {
			str.flagIf = true // true == подтверждение пустой точки
			i = 5       	  // чтоб окончился цикл
			f(xo, x, str) 	  // присвоение точке значения 1 или 2
		}
		if !str.flagIf {
			x = rand.Intn(str.rx*str.rx) // сразу даёт рандомное значение
			i = 0             // для повторения цикла
		}
	}
}

//f == присвоение точки на поле значение X или O
func f(xo []int, x int, str strData) {
	if str.flag { //ход X
		xo[x] = 1
	}
	if !str.flag { //ход O
		xo[x] = 2
	}
}

//win == проверка совпадений линий
func win(xo []int, str *strData)int{
	for i := 1; i <= 2; i++ {
		str.unit = i
		str.flagWin = winDiag1(*str, xo) + winDiag2(*str, xo) + winVertic(*str, xo) + winGoriz(*str, xo)
		if i == 1 {
			if str.flagWin >= 1 {
				winView(str)
				str.xWin++
				i = 3
				return 1
			}
		}
		if i == 2 {
			if str.flagWin >= 1 {
				winView(str)
				str.oWin++
				i = 3
				return 1
			}
		}
	}
	return 0
}

// winView == показ победы
func winView(str *strData){
	if str.unit == 1 {
		fmt.Println("Победа X")
	}
	if str.unit == 2 {
		fmt.Println("Победа O")
	}
}

// функции сравнения по всем линиям
// winGoriz == выполнение сравнений по горизонтали на выйгрыш
func winGoriz(str strData, xo []int)int{
	a := 0
	for i := 0; i <= str.rx-1; i++ {
		for k := 0; k <= str.rx-1; k++ {
			if xo[a] == str.unit {str.check1++}
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
			if xo[a] == str.unit {str.check1++}
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
		if xo[a] == str.unit {str.check1++}

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
		if xo[a] == str.unit {str.check1++}
		a = a+str.rx-1
	}
	if str.check1 == str.rx {
		return 1
	}
	str.check1 = 0
	return 0
}