package main

import "fmt"

type MyRing struct {
	prve  *MyRing
	next  *MyRing
	Value int
}

func NewRing(val ...int) *MyRing {
	if len(val) == 0 {
		return nil
	}

	r := &MyRing{
		Value: val[0],
	}

	p := r
	for i := 1; i < len(val); i++ {
		p.next = &MyRing{
			Value: val[i],
			prve:  p,
		}
		p = p.next
	}
	p.next = r
	r.prve = p
	return r
}

func (r *MyRing) Print() {
	fmt.Print("[")
	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%d", r.Value)
		if i != r.Len()-1 {
			fmt.Printf(", ")
		}
		r = r.Next()
	}
	fmt.Println("]")
}

func (r *MyRing) Len() int {
	if r.next == nil {
		return 1
	}
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
		r.next = &MyRing{}
	}
	return r.next
}

func (r *MyRing) Del() *MyRing {
	if r.Len() <= 1 {
		return nil
	}
	if r.Len() == 2 {
		r.prve.prve = nil
		r.prve.next = nil
		return r.prve
	}
	r.next.prve = r.prve
	r.prve.next = r.next
	return r.prve
}
