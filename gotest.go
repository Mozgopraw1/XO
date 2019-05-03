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
		xo := make([]int, 9)
		// REVIEW: здесь у тебя flag, flagIf всегда == true, зачем их передавать переменными?

		// я не уверен, но походу это влияло на процент побед, но я не уверен, поставил указатели.
		xod(xo, &flag, &flagIf, &flagWin, &NoWin, &Xwin, &Owin)
	}

	fmt.Println("Общий результат:")
	fmt.Println("X:     ", Xwin)
	fmt.Println("O:     ", Owin)
	fmt.Println("Ничья: ", NoWin)
	dt1 := time.Now()
	fmt.Println("Затраченное время на вычисление: ", dt1.Sub(dt))
}

//основная функция по запуску игры
// REVIEW: "`xod` - `x` is unused", означает, что при передаче
//  параметра x в функцию ты его должен там использовать, а иначе
//  тебе незачем его туда передавать. А ты его передаёшь и тут же
//  переназначаешь, а это значит что передавать его нет смысла.
//  Теперь понятно?
// REVIEW: когда в функцию передаётся столько параметров, это
//  убивает возможность сделать тест для функции; это ухудшает
//  читаемость и (!!) контроль над кодом.
func xod(xo []int, flag *bool, flagIf *bool, flagWin *int, NoWin *int, Xwin *int, Owin *int) {
	for i := 0; i <= 8; i++ {
		// REVIEW: вот переназначение x; зачем ты
		// убрал X сверху

		//X == координата точки на поле
		x := rand.Intn(8) // рандомная цифра от 0 до 8

		xoVariant(x, xo, *flagIf, *flag) // проверка хода
		viewXo(xo)                     // показ поля
		win(xo, flagWin, Xwin, Owin)
		flag1(flag) // смена стороны
		if *flagWin != 0 {
			fmt.Println("Действительно выйграл, выходим")
			i = 9
			*flagWin = 0 //сброс значения
		}
		if *flagWin == 0 && i == 8 {
			fmt.Println("Похоже ни кто не выйграл")
			*NoWin++
		} //проверка на ничью
	}
}

//вывод поля
// REVIEW: опечатка в английском слове в названии функции
// исправил
func viewXo(xo []int) {
	fmt.Println(xo[0], xo[1], xo[2])
	fmt.Println(xo[3], xo[4], xo[5])
	fmt.Println(xo[6], xo[7], xo[8])
	fmt.Println(" ")
}

//смена хода
func flag1(flag *bool) {
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
func xoCor(xo []int, x int, flagIf *bool, flag bool) {
	for i := 0; i <= 2; i++ {
		if xo[x] == 0 {
			*flagIf = true
			i = 5
				f(flag, xo, x, i)
		}
		if !*flagIf {
			x = rand.Intn(9) // сразу даёт рандомное значение
			i = 0
		}
	}
}

func f(flag bool, xo []int, x int, i int) int {
	if flag { //ход X
		xo[x] = 1
	}
	if !flag { //ход O
		xo[x] = 2
	}
	i = 3
	return i
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
	*game = 10000
}
