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

func main() {
	// длина стороны поля
	rx := 3
	type OneSt struct {

	}
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
		xo := make([]int, rx*rx)
		xod(xo, &flag, &flagIf, &flagWin, &NoWin, &Xwin, &Owin, rx)
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
func xod(xo []int, flag *bool, flagIf *bool, flagWin *int, NoWin *int, Xwin *int, Owin *int, rx int) {
	for i := 0; i <= 8; i++ {

		//X == координата точки на поле
		x := rand.Intn(8) // рандомная цифра от 0 до 8

		xoVariant(x, xo, *flagIf, *flag) // проверка хода
		viewXo(xo, rx)                   // показ поля
		win(xo, flagWin, Xwin, Owin)     // проверка по вертикали, горизонтали, диагонали на выйгрышь
		playerMove(flag)                 // смена стороны
		// проверка на ничью, и сброс флага выйгрыша для начала новой партии
		if *flagWin != 0 {
			fmt.Println("Действительно выйграл, выходим")
			i = 9
			*flagWin = 0 //сброс значения
		}
		if *flagWin == 0 && i == 8 {
			fmt.Println("Похоже ни кто не выйграл")
			*NoWin++
		}
	}
}

//viewXo == вывод поля в консоль
func viewXo(xo []int, rx int) {
	t := 0
	for i := 0; i <= rx-1; i++ {
		for k := 0; k <= rx-1; k++ {
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
func xoVariant(x int, xo []int, flagIf bool, flag bool) {
	xoCor(xo, x, &flagIf, flag)
	flagIf = false
}

// проверка был ли уже ход в данной точке
// REVIEW: высокая цикломатика (6, надо 5 и меньше)
// не знаю как цикломатику уменьшить в данном случае, описал подробно каждую строку кода тут
func xoCor(xo []int, x int, flagIf *bool, flag bool) {
	for i := 0; i <= 2; i++ { // повторяет цикл бесконечно пока точка на поле не будет пустой
		if xo[x] == 0 {
			*flagIf = true // true == подтверждение пустой точки
			i = 5          // чтоб окончился цикл
			f(flag, xo, x) // присвоение точке значения 1 или 2
		}
		if !*flagIf {
			x = rand.Intn(9) // сразу даёт рандомное значение
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

func win(xo []int, flagWin *int, Xwin *int, Owin *int) {
	for i := 1; i <= 2; i++ {
		if xo[0]&xo[1]&xo[2] == i {
			*flagWin = i
		}
		if xo[3]&xo[4]&xo[5] == i {
			*flagWin = i
		}
		if xo[6]&xo[7]&xo[8] == i {
			*flagWin = i
		}
		if xo[0]&xo[3]&xo[6] == i {
			*flagWin = i
		}
		if xo[1]&xo[4]&xo[7] == i {
			*flagWin = i
		}
		if xo[2]&xo[5]&xo[8] == i {
			*flagWin = i
		}
		if xo[0]&xo[4]&xo[8] == i {
			*flagWin = i
		}
		if xo[2]&xo[4]&xo[6] == i {
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

// winGorriz == выйгрыш по горизонтали
/*func winGoriz(rx int, xo []int, check1 int) {
	for i := 0; i <= rx; i++ {
		for k := 0; k <= rx; i++ {
			if xo[k] == 1 {check1++}
		}
		check1 = 0
		a := +rx
	}
}
*/

func newGame(game *int) {
	fmt.Println("Приветствую тебя в игре XO, сколько желаешь сыграть партий?")
	//var st string

	// REVIEW: необработанное исключение!

	//перестала работать на моём ПК
	/*fmt.Fscan(os.Stdin, &st)

	i1, err := strconv.Atoi(st)
	if err == nil {
		*game = i1
	}
	*/
	*game = 100
}
