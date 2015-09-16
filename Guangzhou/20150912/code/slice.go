package main

import (
	"encoding/json"
	"fmt"
)

// 1, -1, -2, -3, -1, -2, -3
func slice() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3}
	return append(append(s1[:1], s2...), s1[1:]...)
}

// 1, -1, -2, -3, -4, 2, 3, 4
func slice1() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3, -4}
	return append(append(s1[:1], s2...), s1[1:]...)
}

type A struct {
	D1 int
	D2 int
}

type I struct {
	As interface{}
}

func main() {
	fmt.Printf("func slice=%+v\n", slice())
	fmt.Printf("func slice1=%+v\n", slice1())

	data := []byte(`{"As":[{"D1":1, "D2":2},{"D1":-1, "D2":-2}]}`)
	i := new(I)
	i.As = make([]A, 1)
	// i.As = new([]A)
	json.Unmarshal(data, i)
	_, ok := i.As.([]A)
	fmt.Printf("ok=%v\n", ok)
}
