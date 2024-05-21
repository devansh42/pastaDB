package utils

// This implements a single linked list

type Node struct {
	val  interface{}
	next *Node
}
type List struct {
	length     int
	head, tail *Node
}

func (l *List) Len() int {
	return l.length
}

func (l *List) IsNil() bool {
	return l.length == 0
}

func (l *List) Push(val interface{}) int {
	if l.tail != nil {
		l.tail.next = new(Node)
		l.tail = l.tail.next
		l.tail.val = val
		l.length++
		return l.length
	}
	// empty list
	n := new(Node)
	l.head = n
	l.tail = n
	n.val = val
	l.length++
	return l.length
}

func (l *List) Reset() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *List) Head() *Node {
	return l.head
}

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Val() interface{} {
	return n.val
}
