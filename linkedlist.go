package exercism

import "errors"

// Define List and Node types here.
type List struct {
	FirstN *Node
	LastN  *Node
}

type Node struct {
	Value interface{}
	NextN *Node
	PrevN *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.
func NewList(elements ...interface{}) *List {
	if len(elements) == 0 {
		return &List{}
	}

	first := &Node{
		Value: elements[0],
	}

	prev := first
	for _, element := range elements[1:] {
		current := &Node{
			Value: element,
			PrevN: prev,
		}
		prev.NextN = current
		prev = current
	}

	return &List{
		FirstN: first,
		LastN:  prev,
	}
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.NextN
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.PrevN
}

func (l *List) Unshift(v interface{}) {
	prevFirst := l.FirstN

	newNode := &Node{
		Value: v,
		NextN: prevFirst,
	}

	if prevFirst != nil {
		prevFirst.PrevN = newNode
	} else {
		l.LastN = newNode
	}

	l.FirstN = newNode
}

func (l *List) Push(v interface{}) {
	prevLast := l.LastN

	newNode := &Node{
		Value: v,
		PrevN: prevLast,
	}

	if prevLast != nil {
		prevLast.NextN = newNode
	} else {
		l.FirstN = newNode
	}

	l.LastN = newNode
}

func (l *List) Shift() (interface{}, error) {
	first := l.FirstN

	if first == nil {
		return nil, errors.New("The list is empty")
	}

	if first.NextN == nil {
		l.LastN = nil
	} else {
		first.NextN.PrevN = nil
	}

	l.FirstN = first.NextN
	return first.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	last := l.LastN

	if last == nil {
		return nil, errors.New("The list is empty")
	}

	if last.PrevN == nil {
		l.FirstN = nil
	} else {
		last.PrevN.NextN = nil
	}

	l.LastN = last.PrevN
	return last.Value, nil
}

func (l *List) Reverse() {
	for node := l.FirstN; node != nil; {
		next := node.NextN
		node.NextN = node.PrevN
		node.PrevN = next
		node = next
	}

	first := l.FirstN
	l.FirstN = l.LastN
	l.LastN = first
}

func (l *List) First() *Node {
	return l.FirstN
}

func (l *List) Last() *Node {
	return l.LastN
}
