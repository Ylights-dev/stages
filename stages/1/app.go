package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("something else")
	// var a int8 = -1// -128 to 127
	// var b int16 = 32767
	// var c int32 =300000000
	// var d int64 = -1000000000000000000
	// var e uint8 = 10
	// var f uint16 = 99 // 0 to 65535, 2byte
	// var g uint32 = 9999999// 0 to 4bil
	// var h uint64 = 9999999999999999999// 0 to 18pent

	// var I byte // synonym uint8
	// var j rune // synonim int32
	// var k int// в зависимости от платформы занимает 32 или 64
	// var m uint// в зависимости от платформы занимает 32 или 64
	// var a1 float32 = 1.8 // 1.4*10^35 4 byte
	// var b1 float64 = 1.8 // 1.4*10^35
	// var c1 complex64 = 1+2i
	// var d1 complex128 = 4+90i

	// var e1 bool=true
	// var f1 bool=false

	// var name string = "Artur"
	// var name2 string
	// name1 := "Ivan"
	// fmt.Println(name)
	// fmt.Println(name1)
	// fmt.Println(name2)
	var name, age = "Artur", 32
	var c = fmt.Sprintf("My name is %s. And I'm %d yeas old", name, age)
	fmt.Println(c)

}
