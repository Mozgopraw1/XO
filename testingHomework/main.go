package main

import "fmt"

func main() {
	fmt.Println("test")
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

func IntSliceToString([]int{17, 23, 100500}) string{
	return
}