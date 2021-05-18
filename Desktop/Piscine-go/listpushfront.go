package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail.Next = n
		l.Tail = l.Tail.Next
	} else {
		l.Head = n
		l.Tail = n
	}
}
