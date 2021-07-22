package linkedqueue

import (
	"errors"
	"fmt"
	"strings"
)

type LinkedQueue struct {
	front *queueNode
	back  *queueNode
}

type queueNode struct {
	nextInLine *queueNode
	data       int
}

func New() *LinkedQueue {
	return &LinkedQueue{}
}

func (q *LinkedQueue) Push(toAdd int) {
	node := &queueNode{data: toAdd}
	if q.back == nil {
		q.back = node
	}
	if q.front != nil {
		q.front.nextInLine = node
	}
	q.front = node
}

func (q *LinkedQueue) PushAll(toAdd ...int) {
	for _, each := range toAdd {
		q.Push(each)
	}
}

func (q *LinkedQueue) Pop() (int, error) {
	if q.back == nil {
		return -1, errors.New("LinkedQueue: Queue is empty")
	}
	ret := q.back.data
	q.back = q.back.nextInLine

	return ret, nil
}

// func (q *LinkedQueue) isEmpty() bool {
// 	return q.back == nil
// }

func (q *LinkedQueue) String() string {
	str := ""
	curr := q.back
	for curr != nil {
		str += fmt.Sprintf("%d\n", curr.data)
		curr = curr.nextInLine
	}
	return reverseByWords(str)
}

func reverseByWords(str string) string {
	words := strings.Fields(str)
	rev := ""
	for i := len(words); i > 0; i-- {
		rev += words[i-1] + "\n"
	}
	return rev
}
