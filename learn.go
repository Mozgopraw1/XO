package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// test()
	// control()
	unicodeLearn()
}
func test() {
	// Эту вот строку не в коем случае не надо вставлять
	// в быструю функцию, иначе из за неё рандомные
	// числа станут одинаковыми
	rand.Seed(time.Now().UTC().UnixNano())
	// перед копирование слайса требуется создавать
	// слайс новый с длиной копироваемого массива и
	// после этого только копировать туда
	slice5 := make([]int, 15, 15)
	for i := 0; i <= 14; i++ {
		x := rand.Intn(9)
		slice5[i] = x
	}
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	slice8 := append(slice7[:2], slice7[10:]...)
	fmt.Println("из кусков слайса", slice8)

	//ссылка на массива через слайс
	a := [...]int{5, 6, 7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("слайс из массивы", sl8)

	sl9 := make([]int, 10)
	copy(sl9[3:6], []int{1, 2, 3})
	fmt.Println(sl9)

	// %v - сюда передаётся значение sl9
	fmt.Printf("slice elements is %v\nnewline\n", sl9)

	//MAPS
	//MAPS
	//MAPS
	//MAPS

	var mm map[string]string
	fmt.Println("uninitialized map", mm)
	// panic: assignment to entry in nil map
	// mm["test"] = "ok"

	// полная инициализация
	var mm2 map[string]string = map[string]string{}
	// или так
	mmm2 := map[string]string{}
	mm2["test"] = "ok"
	mmm2["test"] = "ok"
	fmt.Println(mm2)

	// короткая инициализация
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Vasily"
	fmt.Println(mm3)

	// получение значения
	firstName := mm3["firstName"]
	fmt.Println("firstName", firstName, len(firstName))

	// если обратиться к несуществующему ключу - отдасться значение по умолчанию
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	// проверка на то, что значение есть
	// true было добавлено, false не было добавлено.
	mm3["lastName"] = ""
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist: ", ok)

	// только получение признака существования
	_, exist := mm3["firstName"]
	fmt.Println("firstName exist:", exist)

	// удаление значения
	// можно ключ сделать как переменную
	keyF := "firstName1"
	delete(mm3, keyF)
	_, exist = mm3["firstName"]
	fmt.Println("firstName exist:", exist)

	mm4 := mm3
	mm4[keyF] = "test"
	fmt.Println(mm3, mm4)

	//УПРАВЛЯЮЩИЕ СТРУКТРЫ / CONTROL
	//УПРАВЛЯЮЩИЕ СТРУКТРЫ / CONTROL
	//УПРАВЛЯЮЩИЕ СТРУКТРЫ / CONTROL
	//УПРАВЛЯЮЩИЕ СТРУКТРЫ / CONTROL
}
func control() {
	a := true
	if a {
		println("hello world")
	}

	b := 1
	c := 22
	if b == 1 && a || c != 22 {
		println("неявное преобразование ( if b ) не работает")
	}

	mm := map[string]string{
		"firstName": "Vasily",
		"lastName": "Romanov",
	}
	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exist, = ", firstName)
	} else {
		println("no firstName")
	}

	if firstName, ok := mm["firstName"]; !ok {
		println("no firstName")
	} else if firstName == "Vasily" {
		println("firstName is Vasily")
	} else {
		println("firstName is not Vasily")
	}

	//циклы
	for {
		println("бесконечный цикл аналог while(true)")
		break
	}

	sl := []int{3,4,5,6,7,8}
	value := 0
	idx := 0

	//операции по slice
	for idx < 4 {
		if idx < 2 {
			idx ++
			continue
		}
		value = sl[idx]
		idx++
		println("while-style loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		println("range slice by index", idx, sl[idx])
	}

	for _, val := range sl {
		println("range slice by idx-value", val)
	}

	//операции по map
	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key, val", key, val)
	}

	for _, val := range mm {
		println("range map by val", val)
	}

	// операция switch
	mm["firstName"] = "Vasily"
	mm["flag"] = "No"
	switch mm["firstName"] {
	case "Vasily", "Evgeny":
		println("switch - name is Vasily")
		// в отличии от других языков  - не переходим в другой вариант по умолчанию
		mm["firstName"] = "Petr"
		fallthrough
		case "Petr":
			if mm["flag"] == "Ok" {
				break // выходим из switch, чтобы не выполнять переход в другой вариант
			}
			println("switch - name is Petr")
			//переход к слеюущему варианту
			fallthrough // переходим в следующий вариант
	default:
		println("switch - some other name")
		mm["firstName"] = "Vasily"
	}

	// как замена можественным if else
	switch {
	case mm["firstName"] == "Vasily":
		println("switch2 - Vasily")
	case mm["lastName"] == "Romanov":
		println("switch2 - Romanov")
	}

	//выходим из цикла ьудучи внутри switch
	Loop:
		for key, val := range mm{
			println("switch in loop", key, val)
			switch {
			case key == "firstName" && val == "Vasily":
				println("switch - break loop here")
				break Loop
			}
		}
	return
}

func unicodeLearn(){
	var symbol rune = 'a'
	var autoSymbol = 'a'
	unicodeSymbol := '⌘'
	unicodeSymbolByNumber := '\u2318'
	println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)

	str1 := "Привет, Мир!"
	fmt.Println("ru: ", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	str2 := "Тест"
	fmt.Println("cn: ", str2, len(str2))
	for index, runeValue := range str2 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}
	println(str2[1])

	bin := []byte(str2)
	fmt.Println("binary cn: ", bin, len(bin))
	for idx, val := range bin {
		fmt.Printf("raw binary idx: %v, oct: %v, hex: %x\n", idx, val, val)
	}

	return
}
