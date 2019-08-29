package main

import (
	"strconv"
)

func main() {
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
		result = result+strconv.Itoa(td[i])
	}
	return result
}

func MergeSlices(sl1[]float32, sl2[]int32) []int{
	var result []int
	len1 := len(sl1)
	len2 := len(sl2)
	for i:=0; i<=len1-1; i++{
		result = append(result, (int(sl1[i])))
	}
	for i:=0; i<=len2-1; i++{
		result = append(result, (int(sl2[i])))
	}
	return result
}

func GetMapValuesSortedByKey(map[int]string) {
	for i:=0;i<11;i++{
		map[int]string
	}
	return result[]
}