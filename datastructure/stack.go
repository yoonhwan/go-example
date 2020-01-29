package container

import (
	"fmt"

	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

type iStack interface {
	push(value int) *DNode
	pop() *DNode
	clear()
	flatmap(fn func(int))
	count() int
}

type Stack struct {
	linkedList *LinkedList
}

func (this *Stack) initial() {
	this.linkedList = &LinkedList{}
	this.linkedList.InitialDoublyLinkedList()
}

func (this *Stack) push(value int) *DNode {
	return this.linkedList.InsertLast(value)
}

func (this *Stack) pop() *DNode {
	return this.linkedList.DeleteLast()
}

func (this *Stack) clear() {
	this.linkedList.Clear()
}

func (this *Stack) flatmap(fn func(int)) {
	this.linkedList.Tour(fn)
}

func (this *Stack) count() int {
	return this.linkedList.Len()
}

func StartStackTest() {
	// testStack := Stack{}

	var testStack *Stack = &Stack{}
	testStack.initial()

	for index := 0; index < 11; index++ {
		testStack.push(index)	
	}
	
	testStack.flatmap(func(x int) { fmt.Println("Stack Tour:", x) })
	fmt.Println("Stack Length:", testStack.count())
	for index := 0; index < 11; index++ {
		fmt.Println("Stack Pop:", testStack.pop())	
	}
	testStack.push(5)
	testStack.push(6)
	fmt.Println("Stack Length:", testStack.count())
	testStack.pop()
	testStack.flatmap(func(x int) { fmt.Println("Stack Tour:", x) })
	testStack.clear()
	fmt.Println("Stack Length:", testStack.count())
}
