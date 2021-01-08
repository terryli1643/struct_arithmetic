/**
lru: 最近最少访问策略
*/
package lru

type MyList struct {
	cap   int
	Value interface{}
	head  *MyList
	tail  *MyList
	prve  *MyList
	next  *MyList
}

var l *MyList
var h map[interface{}]*MyList

func NewMyList(cap int) *MyList {
	head := &MyList{cap: cap, prve: nil}
	tail := &MyList{cap: cap, next: nil}

	head.next = tail
	tail.prve = head

	return &MyList{
		cap:  cap,
		head: head,
		tail: tail,
	}
}

func (l *MyList) Len() int {
	n := 0
	for h := l.head.next; h != l.tail; h = h.next {
		n++
	}
	return n
}

func (l *MyList) Put(v interface{}) {
	e := &MyList{Value: v, cap: l.cap}
	l.tail.prve.insertAfter(e)
	if l.Len() < l.cap {
		e.removeLRU()
	}
	h[v] = e
}

func (l *MyList) Get(v interface{}) *MyList {
	if myList, ok := h[v]; ok {
		return myList
	}

	return nil
}

func (l *MyList) Next() *MyList {
	return l.next
}

func (l *MyList) removeLRU() {
	l.remove(l.head.next)
}

func (l *MyList) remove(e *MyList) {
	e.prve.next = e.next
	e.next.prve = e.prve
}

func (l *MyList) insertAfter(e *MyList) {
	e.next = l.next
	e.prve = l
	e.next.prve = e
	e.prve.next = e
}

func (l *MyList) moveToTail(e *MyList) {
	l.remove(e)

}
