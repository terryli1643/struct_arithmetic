package main

import (
	"fmt"
)

var data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var start = 3
var step = 4

func main() {
	r := NewRing(data...)
	r.Print()

	for i := 0; i < start-1; i++ {
		r = r.Next()
	}

	fmt.Printf("start at %d \n", r.Value)

	for i := 0; i < len(data)-1; i++ {
		var j = step
		for j > 0 {
			j--
			r = r.Next()
		}
		fmt.Printf("remove %d, ", r.Value)
		r = r.Del()
		r.Print()
	}

	fmt.Printf("result is %d \n", r.Value)
}
