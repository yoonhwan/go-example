package container

import (
	"fmt"

	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

type iQueue interface {
	push(value int) *DNode
	pop() *DNode
	clear()
	flatmap(fn func(int))
	count() int
}

type Queue struct {
	linkedList *LinkedList
}

func (this *Queue) initial() {
	this.linkedList = &LinkedList{}
	this.linkedList.InitialDoublyLinkedList()
}

func (this *Queue) push(value int) *DNode {
	return this.linkedList.InsertLast(value)
}

func (this *Queue) pop() *DNode {
	return this.linkedList.DeleteFirst()
}

func (this *Queue) clear() {
	this.linkedList.Clear()
}

func (this *Queue) flatmap(fn func(int)) {
	this.linkedList.Tour(fn)
}

func (this *Queue) count() int {
	return this.linkedList.Len()
}

func StartQueueTest() {
	var testQueue *Queue = &Queue{}
	testQueue.initial()

	for index := 0; index < 11; index++ {
		testQueue.push(index)
	}

	testQueue.flatmap(func(x int) { fmt.Println("Queue Tour:", x) })
	fmt.Println("Queue Length:", testQueue.count())
	for index := 0; index < 11; index++ {
		fmt.Println("Queue Pop:", testQueue.pop())
	}
	testQueue.push(5)
	testQueue.push(6)
	fmt.Println("Queue Length:", testQueue.count())
	testQueue.pop()
	testQueue.flatmap(func(x int) { fmt.Println("Queue Tour:", x) })
	testQueue.clear()
	fmt.Println("Queue Length:", testQueue.count())
}
