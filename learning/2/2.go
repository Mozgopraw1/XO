package main

import (
	"fmt"
	"time"
)

func main() {
	showMeTheMoney()
	fmt.Println("the stuff is ", stuff)
	myTimer := getTimer()
	defer myTimer()

	f := func() {
		myTimer()
	}

	f()
	formatPrint()
}

func showMeTheMoney()  {
	fmt.Printf("$$$$")
}

// sum - несколько входящих значений
func sum(i int, j int) int {
	return i + j
}

// sumLight - упрощенная запись для нескольких значений одинакового типа
func sumLight(i, j int) int {
	return i + j
}

//sumMore - для получения произвольного списка однотипных значений
// входящие и выходящии данные непонятны
func sumMore(stuff ...int) (res int)  {
	for i := range stuff {
		res +=stuff[i]
	}
	return
}

//sumOnlyNatural - возврат нескольких значений
// непонятна строка return 0, fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
func sumOnlyNatural(stuff ...int) (int, error) {
	res := 0
	for i := range stuff {
		if stuff[i] < 0 {
			return 0, fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, nil
}

//sumNatural2 - возврат именнованных значений
func sumNatural2(stuff ...int) (res int, err error) {
	for i := range stuff {
		if stuff[i] <0 {
			err = fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
			return
		}
		res += stuff[i]
	}
	return res, err
}

var stuff = "not ready"

// init - функция выполняющаяся при запуске программы автоматически
func init() {
	stuff = "ready"
}

// getTimer - функция замыкания
func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Println("Time from start", time.Since(start))
	}
}


// formatPrint - тест по форматированиб спецификаторов
func formatPrint() {
	t1 := true
	t2 := false
	fmt.Printf("тест форматирования boolen t1: %t, t2: = %t \n", t1, t2)

	b1 := 1
	b2 := 53
	fmt.Printf("тест форматирования двоичный вывод b1: %b, b2 = %b \n", b1, b2)

	c1 := 'h'
	c2 := 'д'
	fmt.Printf("тест форматирования символа в числовой код c1 = %c, c2 = %c \n", c1, c2)

	d1 := 10
	d2 := 6483
	fmt.Printf("тест форматирования целых чисел в десятичной системе: d1 = %d, d2 = %d \n", d1, d2)

	o1 := 10
	o2 := 6483
	fmt.Printf("тест форматирования целых чисел в восьмиричную систему: o1 = %o, o2 = %o \n", o1, o2)
}
