package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p
} //так писать смысла нет(обычно)
func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

// func totalPrice(initPrice int) func(int) int {
// 	sum := initPrice
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func doSomething(callback func(int, int) int, s string) int {
// 	result := callback(5, 1)
// 	fmt.Println(s)
// 	fmt.Println(result)
// 	return result
// }

func main() {
	p := Point{X: 1, Y: 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)
	v := &p
	v.movePointPtr(1, 1)
	// 	orderPrice := totalPrice(1) //sum=1
	// 	fmt.Println(orderPrice(1))  //sum 1+1=2
	// 	fmt.Println(orderPrice(1))  //sum2+1=3

	// sumCallback := func(n1, n2 int) int {
	// 	return n1 + n2
	// }

	// result := doSomething(sumCallback, "plus")
	// fmt.Println(result)

	// multipleCallback := func(n1, n2 int) int {
	// 	return n1 * n2
	// }
	// result = doSomething(multipleCallback, "multiple")
	// fmt.Println(result)
}
