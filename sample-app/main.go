package main

import "fmt"

func solution(str, ending string) bool {
	// Your code here!
	if len(str) >= len(ending) {
		return str[len(str)-len(ending):] == ending
	} else {
		return false
	}
}
func main() {
	a := "adc"
	b := "adc"
	fmt.Println(solution(a, b))
	fmt.Println(a[len(a)-len(b):])

}
