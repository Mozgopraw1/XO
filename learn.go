package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// перед копирование слайса требуется создавать
	// слайс новый с длиной копироваемого массива и
	// после этого только копировать туда
	slice5 := make([]int, 15,15)
	for i:=0; i<=14; i++{
		rand.Seed(time.Now().UTC().UnixNano())
		x := rand.Intn(9)
		slice5[i] = x
	}
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	return

}