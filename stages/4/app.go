package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type point struct {
	X int `mapstructure:"xx" `
	Y int `mapstructure:"yy"`
}

func (p point) method() {
	fmt.Println("call me baby")
}

func main() {
	pointMap := map[string]int{
		"xx": 1, //продублировали имя полей структуры
		"yy": 2,
	}

	p1 := point{}
	mapstructure.Decode(pointMap, &p1)
	fmt.Println(p1)
}

// otherPointMap := make(map[int]point)

// pointMap["a"] = point{
// 	X: 1,
// 	Y: 2,
// }
//fmt.Println(pointMap)
// // fmt.Println(pointMap["a"])
//otherPointMap[1] = point{Y: 1, X: 2}
// // fmt.Println(otherPointMap)
// // fmt.Println(otherPointMap[1])

// var otherMorePointMap map[string]point
// if otherMorePointMap == nil {
// 	otherMorePointMap = map[string]point{
// 		"a": {Y: 1, X: 2},
// 	}
// }
// //otherMorePointMap["a"] = point{Y: 1, X: 2}
// fmt.Println(otherMorePointMap["a"])
// otherMorePointMap["a"].method()

// key := "a"
// value, ok := otherMorePointMap[key]
// if ok {
// 	fmt.Printf("key=%s exist in map\n", key)
// 	fmt.Println(value)
// }else{
// 	fmt.Printf("key=%s does exist in map\n", key)
// 	fmt.Println(value)
// }

// arr := []string{"a", "b", "c"}
// for _, l := range arr {
// 	fmt.Println(l)
//  }

// }
