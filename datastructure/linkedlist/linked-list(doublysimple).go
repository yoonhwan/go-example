package linkedlist

import "fmt"

type DNode struct {
	Value int
	prev  *DNode
	next  *DNode
}

func (this *DNode) SetValue(value int) {
	this.Value = value
}

func (this *DNode) GetValue() int {
	return this.Value
}

func (this *DNode) SetPrev(prev *DNode) {
	this.prev = prev
}

func (this *DNode) SetNext(next *DNode) {
	this.next = next
}

func (this *DNode) GetPrev() *DNode {
	return this.prev
}

func (this *DNode) GetNext() *DNode {
	return this.next
}

func (this *DNode) Clear() {
	this.Value = 0
	this.prev = nil
	this.next = nil
}

type LinkedList struct {
	doubly_head *DNode
	doubly_tail *DNode
}

func (this *LinkedList) InitialDoublyLinkedList() {
	this.doubly_head = nil
	this.doubly_tail = nil
	this.doubly_head = this.MakeNewNode(-1)
	this.doubly_tail = this.MakeNewNode(-1)
	(*this.doubly_head).SetNext(this.doubly_tail)
	(*this.doubly_tail).SetPrev(this.doubly_head)
}

func (this *LinkedList) MakeNewNode(value int) *DNode {
	node := &DNode{value, nil, nil}
	return node
}

func (this *LinkedList) First() *DNode {
	return (*this.doubly_head).GetNext()
}

func (this *LinkedList) Last() *DNode {
	return (*this.doubly_tail).GetPrev()
}

func (this *LinkedList) Link(prev *DNode, node *DNode, next *DNode) {
	if node == nil {
		(*prev).SetNext(next)
		(*next).SetPrev(prev)
	} else {
		(*prev).SetNext(node)
		(*node).SetPrev(prev)
		(*node).SetNext(next)
		(*next).SetPrev(node)
	}
}
func (this *LinkedList) InsertEmpty(value int) (*DNode, bool) {
	if this.Len() > 0 {
		return this.MakeNewNode(value), false
	} else {
		var node *DNode = this.MakeNewNode(value)
		this.Link(this.doubly_head, node, this.doubly_tail)
		return node, true
	}

}

func (this *LinkedList) InsertFirst(value int) *DNode {
	var node *DNode
	if v, e := this.InsertEmpty(value); e == false {
		var node *DNode = this.First()
		this.Link(this.doubly_head, v, node)
	}
	return node
}

func (this *LinkedList) InsertLast(value int) *DNode {
	var node *DNode
	if v, e := this.InsertEmpty(value); e == false {
		var node *DNode = this.Last()
		this.Link(node, v, this.doubly_tail)
	}
	return node
}

func (this *LinkedList) DeleteFirst() *DNode {
	var node *DNode = this.First()
	if node == this.doubly_tail {
		return nil
	}

	prev := (*node).GetPrev()
	next := (*node).GetNext()
	this.Link(prev, nil, next)
	remove := node
	(*remove).Clear()

	return node
}

func (this *LinkedList) DeleteLast() *DNode {
	var node *DNode = this.Last()
	if node == this.doubly_tail {
		return nil
	}

	prev := (*node).GetPrev()
	next := (*node).GetNext()
	this.Link(prev, nil, next)
	remove := node
	(*remove).Clear()

	return node
}

func (this *LinkedList) Len() int {
	var node *DNode = this.First()
	var length int = 0
	if node == this.doubly_tail {
		return 0
	}

outer:
	for {
		length++
		if (*node).GetNext() == this.doubly_tail {
			break outer
		} else {
			node = (*node).GetNext()
		}
	}

	return length
}

func (this *LinkedList) Clear() {
	var node *DNode = this.First()
outer:
	for {
		if this.Len() == 0 {
			break outer
		}

		if (*node).GetNext() == this.doubly_tail {
			(*node).Clear()
			this.InitialDoublyLinkedList()
			break outer
		} else {
			prev := (*node).GetPrev()
			next := (*node).GetNext()
			this.Link(prev, nil, next)

			remove := node
			(*remove).Clear()
			node = next
		}
	}
}

func (this *LinkedList) Contains(target int) (DNode, bool) {
	var node *DNode = this.First()
	var result *DNode = nil
	var find bool = false
outer:
	for {
		if (*node).GetValue() == target {
			result = node
			find = true
			break outer
		}

		if (*node).GetNext() == nil {
			break outer
		} else {
			node = (*node).GetNext()
		}
	}

	if result != nil {
		return *result, find
	}
	return DNode{}, false
}

type doubly_excute func(int)

func (this *LinkedList) Tour(fn doubly_excute) {
	var node *DNode = this.First()
outer:
	for {
		fn((*node).GetValue())
		if (*node).GetNext() == this.doubly_tail {
			break outer
		} else {
			node = (*node).GetNext()
		}
	}
}

func Doubly() {
	var testLinkedList *LinkedList = &LinkedList{}
	testLinkedList.InitialDoublyLinkedList()

	for index := 0; index < 11; index++ {
		testLinkedList.InsertFirst(index)
	}

	fmt.Println("Doubly Linked List Length:", testLinkedList.Len())
	testLinkedList.DeleteFirst()
	testLinkedList.Tour(func(value int) { fmt.Println("Doubly Linked List:", value) })
	testLinkedList.Clear()
	fmt.Println("Doubly Linked List Length:", testLinkedList.Len())

	for index := 0; index < 11; index++ {
		testLinkedList.InsertLast(index)
	}
	if v, e := testLinkedList.Contains(4); e == true {
		fmt.Println(v)
	}

	testLinkedList.DeleteLast()
	testLinkedList.Tour(func(value int) { fmt.Println("Doubly Linked List:", value) })
	testLinkedList.Clear()
	fmt.Println("Doubly Linked List Length:", testLinkedList.Len())

}
