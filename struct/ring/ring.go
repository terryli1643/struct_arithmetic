package main

import (
	"errors"
	"math"
)

type MyRing struct {
	cap  int
	data []interface{} //使用数组实现循环队列
	head int
	tail int
}

func NewMyRing(cap int) *MyRing {
	return &MyRing{
		cap:  cap,
		data: make([]interface{}, cap),
		head: 0,
		tail: 0,
	}
}

func (r *MyRing) Enqueue(e interface{}) error {
	if (r.tail+1)%r.cap == r.head { // 队满
		return errors.New("queue is full")
	}
	r.data[r.tail] = e
	r.tail = (r.tail + 1) % r.cap
	return nil
}

func (r *MyRing) Dequeue() (e interface{}) {
	if r.tail == r.head { // 对空
		return nil
	}
	e = r.head
	r.head = (r.head + 1) % r.cap
	return e
}

func (r *MyRing) Len() int {
	if r.tail == r.head {
		return 0
	}
	return int(math.Abs(float64(r.tail - r.head)))
}
