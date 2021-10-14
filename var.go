package main

import (
	"fmt"
	"reflect"
)

var globalObj = 3

var GlobalObj2 = 4

var (
	globalObj3 = 5
	GlobalObj4 = 6
)

const (
	a = 1
	b = iota
	c
	d
	e
)

func main() {
	fmt.Println(a, b, c, d, e)

	fmt.Println(max(1, 232,43))
}

func max(f float64, num1, num2 int) int {
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num1))

	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}