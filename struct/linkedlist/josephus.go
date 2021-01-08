package main

import (
	"container/ring"
	"fmt"
)

var data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var start = 3
var step = 4

type MyRing struct {
	next  *MyRing
	Value int
}

func New(val ...int) *MyRing {
	r := new(MyRing)
	for i := 0; i < len(val); i++ {
		r = r.init()
		r.Value = val[i]
		r = r.next
	}
	return r
}

func (r *MyRing) init() *MyRing {
	n := new(MyRing)
	r.next = n
}

func (r *MyRing) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func (r *MyRing) Next() *MyRing {
	if r.next == nil {
		r.next = r.init()
	}
	return r.next
}

var r = ring.New(len(data))
var e *ring.Ring

func init() {
	for i := 0; i < len(data); i++ {
		r.Value = data[i]
		r = r.Next()
	}
	e = r
	print()
}

func main() {
	for i := 0; i < start-1; i++ {
		r = r.Next()
	}

	fmt.Printf("start at %d \n", r.Value)

	for i := 0; i < len(data)-1; i++ {
		var j = step
		for j > 0 {
			r = r.Next()
			if r.Value.(int) != -1 {
				j--
			}
		}
		fmt.Printf("remove %d, ", r.Value)
		r.Value = -1
		print()
	}

	for i := 0; i < len(data); i++ {
		if r.Value.(int) != -1 {
			fmt.Printf("result is %d \n", r.Value)
			break
		}
		r = r.Next()
	}
}

func print() {
	for i := 0; i < e.Len(); i++ {
		fmt.Printf("%d, ", e.Value)
		e = e.Next()
	}
	fmt.Println()
}
