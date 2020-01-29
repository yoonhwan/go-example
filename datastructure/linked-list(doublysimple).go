package linkedlist

import "fmt"

var (
	doubly_head *DNode
	doubly_tail *DNode
)

type DNode struct {
	Value int
	prev  *DNode
	next  *DNode
}

func (this *DNode) set(value int) {
	this.Value = value
}

func (this *DNode) get() int {
	return this.Value
}

func (this *DNode) setPrev(prev *DNode) {
	this.prev = prev
}

func (this *DNode) setNext(next *DNode) {
	this.next = next
}

func (this *DNode) getPrev() *DNode {
	return this.prev
}

func (this *DNode) getNext() *DNode {
	return this.next
}

func (this *DNode) clear() {
	this.Value = 0
	this.prev = nil
	this.next = nil
}

func D_InitialDoublyLinkedList() {
	doubly_head = nil
	doubly_tail = nil
	doubly_head = D_MakeNewNode(-1)
	doubly_tail = D_MakeNewNode(-1)
	(*doubly_head).setNext(doubly_tail)
	(*doubly_tail).setPrev(doubly_head)
}

func D_MakeNewNode(value int) *DNode {
	node := &DNode{value, nil, nil}
	return node
}

func D_First() *DNode {
	return (*doubly_head).getNext()
}

func D_Last() *DNode {
	return (*doubly_tail).getPrev()
}

func D_Link(prev *DNode, node *DNode, next *DNode) {
	if node == nil {
		(*prev).setNext(next)
		(*next).setPrev(prev)
	} else {
		(*prev).setNext(node)
		(*node).setPrev(prev)
		(*node).setNext(next)
		(*next).setPrev(node)
	}
}
func D_InsertEmpty(value int) (*DNode, bool) {
	if D_Len() > 0 {
		return D_MakeNewNode(value), false
	} else {
		var node *DNode = D_MakeNewNode(value)
		D_Link(doubly_head, node, doubly_tail)
		return node, true
	}

}

func D_InsertFirst(value int) *DNode {
	var node *DNode
	if v, e := D_InsertEmpty(value); e == false {
		var node *DNode = D_First()
		D_Link(doubly_head, v, node)
	}
	return node
}

func D_InsertLast(value int) *DNode {
	var node *DNode
	if v, e := D_InsertEmpty(value); e == false {
		var node *DNode = D_Last()
		D_Link(node, v, doubly_tail)
	}
	return node
}

func D_DeleteFirst() *DNode {
	var node *DNode = D_First()
	if node == doubly_tail {
		return nil
	}

	prev := (*node).getPrev()
	next := (*node).getNext()
	D_Link(prev, nil, next)
	remove := node
	(*remove).clear()

	return node
}

func D_DeleteLast() *DNode {
	var node *DNode = D_Last()
	if node == doubly_tail {
		return nil
	}

	prev := (*node).getPrev()
	next := (*node).getNext()
	D_Link(prev, nil, next)
	remove := node
	(*remove).clear()

	return node
}

func D_Len() int {
	var node *DNode = D_First()
	var length int = 0
	if node == doubly_tail {
		return 0
	}

outer:
	for {
		length++
		if (*node).getNext() == doubly_tail {
			break outer
		} else {
			node = (*node).getNext()
		}
	}

	return length
}

func D_Clear() {
	var node *DNode = D_First()
outer:
	for {
		if D_Len() == 0 {
			break outer
		}

		if (*node).getNext() == doubly_tail {
			(*node).clear()
			D_InitialDoublyLinkedList()
			break outer
		} else {
			prev := (*node).getPrev()
			next := (*node).getNext()
			D_Link(prev, nil, next)

			remove := node
			(*remove).clear()
			node = next
		}
	}
}

func D_Contains(target int) (DNode, bool) {
	var node *DNode = D_First()
	var result *DNode = nil
	var find bool = false
outer:
	for {
		if (*node).get() == target {
			result = node
			find = true
			break outer
		}

		if (*node).getNext() == nil {
			break outer
		} else {
			node = (*node).getNext()
		}
	}

	if result != nil {
		return *result, find
	}
	return DNode{}, false
}

type doubly_excute func(int)

func D_Tour(fn doubly_excute) {
	var node *DNode = D_First()
outer:
	for {
		fn((*node).get())
		if (*node).getNext() == doubly_tail {
			break outer
		} else {
			node = (*node).getNext()
		}
	}
}

func Doubly() {
	D_InitialDoublyLinkedList()

	for index := 0; index < 11; index++ {
		D_InsertFirst(index)
	}

	fmt.Println("Doubly Linked List Length:", D_Len())
	D_DeleteFirst()
	D_Tour(func(value int) { fmt.Println("Doubly Linked List:", value) })
	D_Clear()
	fmt.Println("Doubly Linked List Length:", D_Len())

	for index := 0; index < 11; index++ {
		D_InsertLast(index)
	}
	if v, e := D_Contains(4); e == true {
		fmt.Println(v)
	}

	D_DeleteLast()
	D_Tour(func(value int) { fmt.Println("Doubly Linked List:", value) })
	D_Clear()
	fmt.Println("Doubly Linked List Length:", D_Len())

}
