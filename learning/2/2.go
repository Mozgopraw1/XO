package main

import "fmt"

func main() {
	showMeTheMoney()
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