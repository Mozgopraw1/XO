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
	rx int	// == ширина/длина поля
	check1 int	// == временная перемення в winD,V,G
	unit int // == крестик или нолик (1 или 2)
	winF int // == win передача флага об проверке
}

func main() {
	//инициализация структыр
	str := new(strData)

	// шаирина/длина поля
	str.rx = 100

	// ход игрока; true == x; false == o;
	flag := true

	// флаг свободной точки.
	flagIf := false

	// количество партий которые хочет сыграть игрок
	var game int
	newGame(&game)

	//узнаёт время в данный момент
	dt := time.Now()

	//Включает рандомную функцию
	rand.Seed(time.Now().UTC().UnixNano())

	// 0 ни кто не выйграл, 1 выйграл Х, 2 выйграл О
	var flagWin int

	//сколько побед
	Xwin := 0
	Owin := 0
	NoWin := 0

	//запуск всей игры
	for i := 1; i <= game; i++ {
		xo := make([]int, str.rx*str.rx)
		xod(xo, &flag, &flagIf, &flagWin, &NoWin, &Xwin, &Owin, *str)
	}

	// общая статистика в конце игры
	fmt.Println("Общий результат:")
	fmt.Println("X:     ", Xwin)
	fmt.Println("O:     ", Owin)
	fmt.Println("Ничья: ", NoWin)
	dt1 := time.Now()
	fmt.Println("Затраченное время на вычисление: ", dt1.Sub(dt))
}

//xod == основная функция по запуску игры(буду переделывать, заменяя набор переменных на структуры)
func xod(xo []int, flag *bool, flagIf *bool, flagWin *int, NoWin *int, Xwin *int, Owin *int, str strData) {
	for i := 0; i <= str.rx*str.rx-1; i++ {

		//X == координата точки на поле
		x := rand.Intn(str.rx*str.rx-1) // рандомная цифра от 0 до RX*RX
		xoVariant(x, xo, *flagIf, *flag, str) // проверка хода
		str.winF = win(xo, flagWin, Xwin, Owin, str)     // проверка по вертикали, горизонтали, диагонали на выйгрышь
		if str.winF >= 1 {
			i = str.rx*str.rx+1

			viewXo(xo, str)                   // показ поля
		}
		playerMove(flag)                 // смена стороны
		// проверка на ничью, и сброс флага выйгрыша для начала новой партии
		if *flagWin != 0 {
			fmt.Println("Действительно выйграл, выходим")
			i = str.rx*str.rx+1
		}
		if *flagWin == 0 && i == str.rx*str.rx-1 {

			fmt.Println("Похоже ни кто не выйграл")
			viewXo(xo, str)                   // показ поля
			*NoWin++
		}
		*flagWin = 0 //сброс значения
	}
}

//viewXo == вывод поля в консоль
func viewXo(xo []int, str strData) {
	t := 0
	for i := 0; i <= str.rx-1; i++ {
		for k := 0; k <= str.rx-1; k++ {
			fmt.Print(xo[t], " ")
			t++
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

//playerMove == смена хода
func playerMove(flag *bool) {
	if *flag {
		*flag = false
	} else {
		*flag = true
	}
}

// ход на поле
// проверка чей ход
func xoVariant(x int, xo []int, flagIf bool, flag bool, str strData) {
	xoCor(xo, x, &flagIf, flag, str)
	flagIf = false
}

// проверка был ли уже ход в данной точке
// REVIEW: высокая цикломатика (6, надо 5 и меньше)
// не знаю как цикломатику уменьшить в данном случае, описал подробно каждую строку кода тут
func xoCor(xo []int, x int, flagIf *bool, flag bool, str strData) {
	for i := 0; i <= 2; i++ { // повторяет цикл бесконечно пока точка на поле не будет пустой
		if xo[x] == 0 {
			*flagIf = true // true == подтверждение пустой точки
			i = 5          // чтоб окончился цикл
			f(flag, xo, x) // присвоение точке значения 1 или 2
		}
		if !*flagIf {
			x = rand.Intn(str.rx*str.rx) // сразу даёт рандомное значение
			i = 0            // для повторения цикла
		}
	}
}

//f == присвоение точки на поле значение X или O
func f(flag bool, xo []int, x int) {
	if flag { //ход X
		xo[x] = 1
	}
	if !flag { //ход O
		xo[x] = 2
	}
}

//проверка совпадений линий
// REVIEW: высокая цикломатика (12, надо 5 и меньше)
// REVIEW: зачем Xwin и Owin с заглавной? Названияя
//  параметров должны быть с нижнего регистра и это
//  не единственное место.

func win(xo []int, flagWin *int, Xwin *int, Owin *int, str strData)int{
	for i := 1; i <= 2; i++ {
		str.unit = i
		flagWin := winDiag1(str, xo) + winDiag2(str, xo) + winVertic(str, xo) + winGoriz(str, xo)
		if i == 1 {
			if flagWin >= 1 {
				fmt.Println("Победа ", str.unit)
				*Xwin++
				i = 3
				return 1
			}
		}
		if i == 2 {
			if flagWin >= 1 {
				fmt.Println("Победа ", str.unit)
				*Owin++
				i = 3
				return 1
			}
		}
	}
	return 0
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

func newGame(game *int) {
	//var st string

	// REVIEW: необработанное исключение!

	//перестала работать на моём ПК
	/*fmt.Fscan(os.Stdin, &st)

	i1, err := strconv.Atoi(st)
	if err == nil {
		*game = i1
	}
	*/
	*game = 2
}
