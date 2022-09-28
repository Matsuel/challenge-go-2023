package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = val
		l.Head = val
	} else {
		l.Head = val
		l.Head.Next = val
	}
}
