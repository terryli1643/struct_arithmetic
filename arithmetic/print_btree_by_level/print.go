package main

import (
	"container/list"
	"fmt"
	"time"
)

// 按照层级依次打印一个二叉树（利用队列）
func main() {
	l := list.New()
	l.PushFront(newBTree())
	print(l)
}

func print(l *list.List) {
	for e := l.Back(); e != nil; e = l.Back() {
		t := e.Value.(*BTree)
		if t.left != nil {
			l.PushFront(t.left)
		}
		if t.right != nil {
			l.PushFront(t.right)
		}

		l.Remove(e)
		fmt.Println(t.value)
		time.Sleep(1 * time.Second)
	}
}

type BTree struct {
	left  *BTree
	right *BTree
	value int
}

func newBTree() *BTree {
	root := &BTree{value: 0}
	n1 := &BTree{value: 1}
	n2 := &BTree{value: 2}
	n3 := &BTree{value: 3}
	n4 := &BTree{value: 4}
	n5 := &BTree{value: 5}
	n6 := &BTree{value: 6}
	n7 := &BTree{value: 7}
	n8 := &BTree{value: 8}
	n9 := &BTree{value: 9}
	n10 := &BTree{value: 10}
	n11 := &BTree{value: 11}
	n12 := &BTree{value: 12}
	root.left = n1
	root.right = n2
	n1.left = n3
	n2.right = n4
	n3.left = n5
	n3.right = n6
	n4.left = n7
	n4.right = n8
	n5.left = n9
	n5.right = n10
	n6.left = n11
	n6.right = n12
	return root
}
