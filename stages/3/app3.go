package main

// type Point struct {
// 	X int
// 	Y int
// 	S string
// }

// func (p Point) method() {
// 	fmt.Println(p.X)
// 	fmt.Println(p.Y)
// 	fmt.Println(p.S)
// }

func main() {

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// //lt:=s[5:6]// c пятого по шестой не включительно = только "6"
	// t := s[:5] // тоже самое t := s[0:5]
	// fmt.Println(t)
	// tt := s[5:] // тоже самое   tt := s[5:10]
	// fmt.Println(tt)

	// 	animalsArr := [4]string{
	// 		"dog",
	// 		"cat",
	// 		"giraffe",
	// 		"elephant",
	// 	}
	// 	// animalSlic:=[]string{
	// 	// 	"dog",
	// 	// 	"cat",
	// 	// 	"giraffe",
	// 	// 	"elephant",
	// 	// }
	// 	var a []string = animalsArr[0:2] // от нулевого до второго
	// 	b := animalsArr[1:3]             // от первого(положение) до третьего по порядку(пересчет)
	// 	fmt.Println(a)
	// 	fmt.Println(b)

	// b[0]="123" // изменится массив и слайс соответственно

	//slices - слайсы, кусочки массивов, не хранят в себе данные
	// letters := []string{"a", "b", "c"}
	// letters[0] = "1"
	// letters = append(letters, "d")
	// letters = append(letters, "e", "f")
	// fmt.Println(letters)

	// createSlice := make([]string, 3)
	// createSlice[0] = "1"
	// createSlice[1] = "2"
	// createSlice[2] = "3"
	// fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	// fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	// createSlice = append(createSlice, "4")
	// createSlice[3] = "4"
	// fmt.Println(createSlice)
	// fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	// fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	// createSlice = append(createSlice, "4")
	// fmt.Println(createSlice)
	// fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	// fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	// createSlice = append(createSlice, "4")
	// fmt.Println(createSlice)
	// fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	// fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	// createSlice = append(createSlice, "4")
	// fmt.Println(createSlice)
	// fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	// fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	// // p1 := Point{
	// 	X: 1,
	// 	Y: 2,
	// 	S: "hello",
	// }
	// p2 := &p1
	// p2.method()
	// var a [2]string
	// a[0] = "hello"
	// a[1] = "world"
	// fmt.Println(a)
	// fmt.Println(a[1])

	// numbers:=[...]int{1, 2, 3} //компилятор посчитай мне кол-во эл-ов
	// numbers[2]=4
}

// func structs() {
// 	p1 := Point{
// 		X: 1,
// 		Y: 2,
// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p1.X)
// 	p1.X = 123
// 	fmt.Println(p1)
// 	p2 := Point{
// 		X: 123,
// 	}
// 	fmt.Println(p2)

// 	g := &p1
// 	fmt.Println(g.X)
// 	fmt.Println(g)
// 	fmt.Println(&g)
// }

// func pointers() { //указатели содержат ссылку на адрес памяти где записано значение
// 	a := "hello world" //го занимает память по определенному адресу и пишет туда значение
// 	// b := 42
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	p := &a
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	*p = "oh my"
// 	fmt.Println(a)
// 	b := 32
// 	g := &b
// 	*g = *g / 2
// 	fmt.Println(b)
// }
