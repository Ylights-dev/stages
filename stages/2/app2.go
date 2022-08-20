package main

import (
	"fmt"
)

// const a string = "name"

func main() {
	// stack LIFO - Last In First Out последний дефер выполниться первым
	defer fmt.Println("world") // отложенное выполнение операций после всего остального внутри данного метода
	defer fmt.Println("1")     // отложенное выполнение операций после всего остального внутри данного метода
	defer fmt.Println("2")     // отложенное выполнение операций после всего остального внутри данного метода
	defer fmt.Println("3")     // отложенное выполнение операций после всего остального внутри данного метода
	fmt.Println("hello")

	// x, x1, x2 := test()
	// x3 := test1()
	// fmt.Println(x, x1, x2, x3)
	// 	sum:=0
	// 	for i:=0; i<10; i++ // тоже самое что i+=1 или i=i+1
	// 	{
	// sum+=i
	// 	}
	// 	for ;sum<1000;{ // while, until без точей с запятой
	// 		sum+=10
	// 	}
	// a := 0
	// for a < 1000 {
	// 	//i := isTest(a)
	// 	switch isTest(a) {
	// 	case 1:
	// 		fmt.Println("a=2")
	// 	case 2:
	// 		fmt.Println("a=3")
	// 	default:
	// 		fmt.Printf("unknown a, a=%d\n", a)
	// 	}
	// 	// if i == 1 {
	// 	// 	fmt.Println("a=2")
	// 	// } else if i == 2 {
	// 	// 	fmt.Println("a=3")
	// 	// } else {
	// 	// 	fmt.Printf("unknown a, a=%d\n", a)
	// 	// }
	// 	a++
	// }
	// //////////////////////////////////////////////////////////////////////////////////////////////
	// a := 2
	// i := isTest(2)
	// switch {
	// case i == 1:
	// 	fmt.Println("a=2")
	// case i == 2:
	// 	fmt.Println("a=3")
	// default:
	// 	fmt.Printf("unknown a, a=%d\n", a)
	// }
}

// /////////////////////////////////////////////////////////////////////////////
func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}

//////////////////////////////////////////////////////////////////////////////
// type s struct {
// 	a string
// 	b string
// 	c string
// }

// func test() (a, b, c string) {
// 	a = "Hello"
// 	c = "all"
// 	b = "world"
// 	return a, b, c
// }

// func test1() string {
// 	return "empty"
// }
