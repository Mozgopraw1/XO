package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("test")
	td := make([]int, 2)
	td[0] = 16
	td[1] = 256
	td[2] = 24
	var test2 string
	test2 = IntSliceToString(td)
	fmt.Println(test2)
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int{
	return []int{1, 2, 3}
}

func IntSliceToString(td[]int) string{
	var result string
	for i:=0; i<=2; i++ {
		result = strconv.Itoa(td[i])
	}
	return result
}