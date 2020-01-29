package linkedlist

import "fmt"

var (
	head *SNode
)

type SNode struct {
	Value int
	next  *SNode
}

func (this *SNode) set(value int) {
	this.Value = value
}

func (this *SNode) get() int {
	return this.Value
}

func (this *SNode) setNext(node *SNode) *SNode {
	this.next = node
	return node
}

func (this *SNode) getNext() *SNode {
	return this.next
}

func (this *SNode) clear() {
	this.Value = 0
	this.next = nil
}

func MakeNewNode(value int) *SNode {
	node := new(SNode)
	(*node).set(value)
	return node
}

func MakeNext(value int) *SNode {
	if head == nil {
		panic("")
	}

	var node *SNode = head
outer:
	for {
		if (*node).getNext() == nil {
			node = (*node).setNext(MakeNewNode(value))
			break outer
		} else {
			node = (*node).getNext()
		}
	}

	return node
}

type excute func(int)

func Tour(fn excute) {
	if head == nil {
		panic("")
	}

	var node *SNode = head
outer:
	for {
		fn((*node).get())
		if (*node).getNext() == nil {
			break outer
		} else {
			node = (*node).getNext()
		}
	}
}

func Len() int {
	if head == nil {
		panic("")
	}

	var node *SNode = head
	var length int = 0
outer:
	for {
		if node != nil {
			length++
		}
		if (*node).getNext() == nil {
			break outer
		} else {
			node = (*node).getNext()
		}
	}

	return length
}

func Clear() {
	if head == nil {
		panic("")
	}

	var node *SNode = head
	var prev *SNode = nil
outer:
	for {
		if Len() == 1 {
			break outer
		}

		if (*node).getNext() == nil {
			prev.clear()
			node = head
		} else {
			prev = node
			node = (*node).getNext()
		}
	}
}

func InsertPrev(target int, value int) {
	if head == nil {
		panic("")
	}

	var node *SNode = head
	var prev *SNode = nil
outer:
	for {
		if (*node).get() == target {
			if prev == nil { //head
				newone := MakeNewNode(value)
				(*newone).setNext(node)
				head = newone
			} else {
				newone := MakeNewNode(value)
				(*newone).setNext(node)
				(*prev).setNext(newone)
			}

			break outer
		}

		if (*node).getNext() == nil {
			break outer
		} else {
			prev = node
			node = (*node).getNext()
		}
	}
}

func InsertNext(target int, value int) {
	if head == nil {
		panic("")
	}

	var node *SNode = head
outer:
	for {
		if (*node).get() == target {
			newone := MakeNewNode(value)
			(*newone).setNext((*node).getNext())
			(*node).setNext(newone)
			break outer
		}

		if (*node).getNext() == nil {
			break outer
		} else {
			node = (*node).getNext()
		}
	}
}

func Delete(target int) {
	if head == nil {
		panic("")
	}

	var node *SNode = head
	var prev *SNode = nil
outer:
	for {
		if (*node).get() == target {
			(*prev).setNext((*node).getNext())
			(*node).clear()
			break outer
		}

		if (*node).getNext() == nil {
			break outer
		} else {
			prev = node
			node = (*node).getNext()
		}
	}
}

func Contains(target int) (SNode, bool) {
	if head == nil {
		panic("")
	}

	var node *SNode = head
	var result *SNode = nil
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
	return SNode{}, false
}

func Singly() {
	head = MakeNewNode(0)
	for index := 1; index < 11; index++ {
		MakeNext(index)
	}

	Tour(func(x int) { fmt.Println(x) })
	InsertPrev(0, -99)
	InsertNext(3, -152125)
	InsertPrev(-152125, -9977)
	if v, e := Contains(4); e == true {
		fmt.Println(v)
	}
	Delete(4)
	if v, e := Contains(4); e == true {
		fmt.Println(v)
	}
	Tour(func(x int) { fmt.Println(x) })
	Clear()
	head.clear()
	head = nil
}
