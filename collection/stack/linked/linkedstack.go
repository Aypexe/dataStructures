package linkedstack

import (
	"errors"
	"fmt"
)

type LinkedStack struct {
	top *stackNode
}

type stackNode struct {
	nextInLine *stackNode
	data       int
}

func New() *LinkedStack {
	return &LinkedStack{}
}

func (q *LinkedStack) Push(toAdd int) {
	node := &stackNode{data: toAdd}
	if q.top != nil {
		node.nextInLine = q.top
	}
	q.top = node
}

func (q *LinkedStack) PushAll(toAdd ...int) {
	for _, each := range toAdd {
		q.Push(each)
	}
}

func (q *LinkedStack) Pop() (int, error) {
	if q.top == nil {
		return -1, errors.New("LinkedStack: Stack is empty")
	}
	ret := q.top.data
	q.top = q.top.nextInLine

	return ret, nil
}

// func (q *LinkedStack) isEmpty() bool {
// 	return q.back == nil
// }

func (q *LinkedStack) String() string {
	str := ""
	curr := q.top
	for curr != nil {
		str += fmt.Sprintf("%d\n", curr.data)
		curr = curr.nextInLine
	}
	return str
}
