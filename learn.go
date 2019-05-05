package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// Эту вот строку не в коем случае не надо вставлять
	// в быструю функцию, иначе из за неё рандомные
	// числа станут одинаковыми
	rand.Seed(time.Now().UTC().UnixNano())
	// перед копирование слайса требуется создавать
	// слайс новый с длиной копироваемого массива и
	// после этого только копировать туда
	slice5 := make([]int, 15,15)
	for i:=0; i<=14; i++{
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
	a := [...]int{5,6,7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("слайс из массивы", sl8)
	return

}